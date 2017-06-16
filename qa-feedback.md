# QA Feedback

Disclaimer: Ik heb geen ervaring met Go :)

### HIGH

- Betekenisvolle automatische unit en integratie tests ontbreken.

- Korte naamgeving. Het is niet altijd direct duidelijk wat een variable betekent.
    - `r` als render, `db`, `mgo` ([building_controller.go](https://github.com/Badmuts/hsleiden-ipsenh-api/blob/master/controller/building_controller.go))
    - `err` als error, `er` als andere error, `o` als sensor, `r` als room ([datapoint_controller.go:55](https://github.com/Badmuts/hsleiden-ipsenh-api/blob/master/controller/datapoint_controller.go#L55))

- Splits grote methoden op in kleinere methoden om de leesbaarheid en testbaarheid te vergroten. ([datapoint_controller:52](https://github.com/Badmuts/hsleiden-ipsenh-api/blob/master/controller/datapoint_controller.go#L52))

- Separation of concern. Door verschillende verantwoordelijkheiden te verdelen over losse componenten kun je leesbaarheid en testbaarheid van je code vergroten.
  Bijv. [building_controller.go](https://github.com/Badmuts/hsleiden-ipsenh-api/blob/master/controller/building_controller.go) bevat nu logica om requests en responses te lezen/genereren, modellen te persisteren (serializatie en database operaties). Het zou interessant kunnen zijn om dat te verdelen over services.

### MEDIUM

- Voorkom code duplicatie.
  - [deploy.js](https://github.com/Badmuts/hsleiden-ipsenh-api/blob/master/operations/scripts/deploy.js) lijkt veel overeen te komen met [die van het front project](https://github.com/Badmuts/hsleiden-ipsenh-front/blob/master/operations/scripts/deploy/index.js).
  - Code duplicatie geld niet alleen voor grote blokken. Bijv. voor het ophalen van een ID uit het request (`bson.ObjectIdHex(mux.Vars(req)["id"])`) zou je een helper functie kunnen maken zodat de exacte manier van ophalen maar op een plek staat.

### LOW

- Verwijder code ipv het uit te commenten ([db.go](https://github.com/Badmuts/hsleiden-ipsenh-api/blob/master/db/db.go))

