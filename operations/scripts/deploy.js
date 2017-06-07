#!/usr/bin/env node
/**
 * usage: ./operations/scripts/deploy.js --branch BRANCH --tag IMAGE_TAG --env ENV
 * 
 *      --branch    Branch that is deployed
 *      --tag       image_TAG to deploy
 *      --env       environment to deploy to <testing|staging|production>
 * 
 */
var fs = require('fs');
var sh = require('shelljs');
var Promise = require('bluebird');
var axios = require('axios');
var chalk = require('chalk');
var yargs = require('yargs').argv;
var log = console.log;

var awsInstance = 'ec2-35-158-19-66.eu-central-1.compute.amazonaws.com';
var travisBranch = yargs.branch;
var current = yargs.tag;
var env = yargs.env;

// @todo Check testing, staging, production
if (env != 'testing') {
    return log(chalk.cyan('INFO ::'), 'Deployment is disabled for non testing environments');
}

var compose = 'docker-compose.' + env + '.yml';

// Check if compose file exists
if (!fs.existsSync(compose)) {
    return log(chalk.red('ERROR ::'), 'No such file:', compose)
}

if (!current) {
    current = sh.exec("echo `echo " + travisBranch + " | cut -d'/' -f 2-`-$(git rev-parse HEAD | cut -c1-7)", {silent:true}).stdout.replace('\n', '');
}

var repoOwner = 'badmuts';
var repoSlug = `${repoOwner}/hsleiden-ipsenh-api`;
var githubEndpoint = 'https://api.github.com';
var token = process.env.IPSENH_GITHUB_TOKEN;
var github = axios.create({
    baseURL: githubEndpoint,
    headers: {
        'Authorization' : `token ${token}`,
        'Content-Type': 'application/json'
    }
})

var slackHookEndpoint = 'https://hooks.slack.com/services/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';
var slack = (data) => axios.post(slackHookEndpoint, data);

// @todo Check for PRs/pushes
// @todo Build deployment url

/**
 * Create a Github deployment for this push
 */
function createDeployment() {
    return github.post(`/repos/${repoSlug}/deployments`, {
        ref: process.env.TRAVIS_COMMIT || travisBranch,
        description: `Deploying $TRAVIS_BRANCH to ${env}`,
        environment: env,
    })
    .catch(err  => log(chalk.red('COULD NOT CREATE DEPLOYMENT'), err));
}

/**
 * Update deployment status with given data.
 * 
 * @param {int} id id of deployment to update
 * @param {Object} data Data to post
 */
function updateDeployment(id, data) {
    return github.post(`/repos/${repoSlug}/deployments/${id}/statuses`, data, {
        headers: {
            'Accept': 'application/vnd.github.ant-man-preview+json'
        }
    })
    .catch(err => log(chalk.red('COULD NOT UPDATE DEPLOYMENT STATUS'), err))
}

// create deployment
log(chalk.cyan('CREATING DEPLOYMENT...'));
createDeployment()
.then(res => {
    log(chalk.green('CREATED DEPLOYMENT'));
    return res.data;
})
.then(deployment => {
    // time to deploy
    // Read compose file into cat to later output on server
    var composeProjectName = `api${travisBranch}`;
    sh.cat(compose)
    .exec(`ssh -o "IdentitiesOnly yes" \
            -o "StrictHostKeyChecking no" \
            -o "User ec2-user" \
            -i "$(pwd)/operations/secrets/travis-aws" \
            ${awsInstance} "mkdir .deploy/; \
            cat > .deploy/docker-compose.yml; \
            TRAVIS_BRANCH=${travisBranch} CURRENT=${current} docker-compose -p ${composeProjectName} -f .deploy/docker-compose.yml up -d; \
            rm -rf .deploy/"`, 
    {silent:true}, // Default behaviour prints everything to console
    function(code, stdout, stderr) {
        var logUrl = `https://travis-ci.com/Badmuts/hsleiden-ipsenh-api/builds/${process.env.TRAVIS_BUILD_ID}`;

        if (code > 0) {
            log(chalk.red('ERROR :: '), stderr)
            // Deployment failed, update status.
            return updateDeployment(deployment.id, {
                state: 'failure',
                description: `Could not deploy ${travisBranch} ${stderr}`,
                log_url: logUrl
            })
            .then(function() {
                slack({
                    "channel": "#builds",
                    "username": "Deploy Bot",
                    "icon_emoji": ":unicorn_face:",
                    "attachments": [
                        {
                            "fallback": `https://github.com/${repoSlug} deploy failed: ${current}`,
                            "pretext": `<https://github.com/${repoSlug}|${repoSlug}> deploy failed: :cd: ${current}`,
                            "color": "bad",
                            "fields": [{
                                "title": "Image :cd:",
                                "value": `<https://hub.docker.com/r/badmuts/hsleiden-ipsenh-api/tags/|${current}>`,
                                "short": true
                            }, {
                                "title": "Environment",
                                "value": env,
                                "short": true
                            }, {
                                "title": "Deploy url",
                                "value": `<${envUrl}|${travisBranch}.api.ipsenh.daan.codes>`,
                                "short": true
                            }, {
                                "title": "Travis build :construction_worker:",
                                "value": `<${logUrl}|#${process.env.TRAVIS_BUILD_NUMBER}>`,
                                "short": true
                            }]
                        }
                    ]
                })
                .catch(err => log(chalk.red('COULD NOT NOTIFY SLACK'), err))
            })
            .catch(err => log(chalk.red('COULD NOT UPDATE DEPLOYMENT STATUS'), err))
        }

        log(chalk.green('SUCCESSFUL DEPLOY'))
        log(chalk.cyan('UPDATING GITHUB DEPLOYMENT...'))

        // Update deployment
        var envUrl = `https://${travisBranch}.api.ipsenh.daan.codes`;
        updateDeployment(deployment.id, {
            state: 'success',
            environment_url: envUrl,
            description: `Branch ${travisBranch} deployed to ${envUrl}`,
            log_url: logUrl
        })
        .then(res => {
            log(chalk.green('GITHUB DEPLOYMENT UPDATED'))
            log(chalk.cyan('SENDING SLACK NOTIFICATION...'))

            slack({
                "channel": "#builds",
                "username": "Deploy Bot",
                "icon_emoji": ":unicorn_face:",
                "attachments": [
                    {
                        "fallback": `https://github.com/${repoSlug} deployed successfully: ${current}`,
                        "pretext": `<https://github.com/${repoSlug}|${repoSlug}> deployed successfully: :cd: ${current}`,
                        "color": "good",
                        "fields": [{
                            "title": "Image :cd:",
                            "value": `<https://hub.docker.com/r/badmuts/hsleiden-ipsenh-api/tags/|${current}>`,
                            "short": true
                        }, {
                            "title": "Environment",
                            "value": env,
                            "short": true
                        }, {
                            "title": "Deploy url",
                            "value": `<${envUrl}|${travisBranch}.api.ipsenh.daan.codes>`,
                            "short": true
                        }, {
                            "title": "Travis build :construction_worker:",
                            "value": `<${logUrl}|#${process.env.TRAVIS_BUILD_NUMBER}>`,
                            "short": true
                        }]
                    }
                ]
            })
            .then(res => log(chalk.green('NOTIFIED SLACK')))
            .catch(err => log(chalk.red('ERROR NOTIFYING SLACK'), err))
        })
        .catch(err => log(chalk.red('COULD NOT UPDATE DEPLOYMENT STATUS'), err));
    });
});