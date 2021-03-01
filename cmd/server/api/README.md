# API DOC

## JOBs
### Create JOB

Create job executes the order for search in ads library

**URL**

`/api/facebook/ads/create_job`

**TYPE**

`POST`

**HEADER**

```
Content-Type : application/json
```

**BODY**

```json
{
    "search_terms": "Cochabamba",
    "access_token": "FACEBOOK TOKEN",
    "page_total": 1000,
    "search_total": 5000,
    "ad_active_status": "ALL",
    "ad_delivery_date_max": "2021-02-03",
    "ad_delivery_date_min": "2021-01-01",
    "ad_reached_countries": "BO"

}
```

**RESPONSE**

```json

{
    "search_terms": "José Carlos Sánchez Verazaín",
    "job_id": 1,
    "access_token": "EAAG.....2tDrPTs7nH10Jf",
    "total_ads": 48
}

```



### List JOBS

This list all the cread jobs


**URL**

`/api/facebook/jobs/list/all`

**TYPE**

`GET`

**URL PARAMS**

* `page_location: int`
* `page_size: int`

**Example**
`/api/facebook/jobs/list/all?page_location=1&page_size=10`


**RESPONSE**

```json

[
    {
        "facebook_job": {
            "search_terms": "José Carlos Sánchez Verazaín",
            "access_token": "EAAGOv4XxRSMBANOjt......7nH10Jf",
            "page_total": 1000,
            "search_total": 5000,
            "ad_active_status": "ALL",
            "ad_delivery_date_max": "2021-02-28",
            "ad_delivery_date_min": "2021-01-01",
            "ad_reached_countries": "BO"
        },
        "total_found_ads": 49,
        "job_id": 1
    }
]
```



## Facebook Ads

Once you have created Jobs, you can list Facebook ads.
### List all facebook ads.

It uses pagination for show the acutal facebook ads.

**URL**

`/api/facebook/ads/list/all/?page_location=1&page_size=10`

**TYPE**

`GET`

**URL PARAMS**

* `page_location: int`
* `page_size: int`


**RESPONSE**

A list of ads.

```json
[
    {
        "ad_id": "897368761011431",
        "page_id": "101029268584854",
        "page_name": "José Carlos Sánchez Verazaín",
        "ad_snapshot_url": "https://www.facebook.com/ads/archive/render_ad/?id=897368761011431&access_token=EA......0Jf",
        "ad_creative_body": "DESPIERTA COCHABAMBA! Es momento de echar a los corruptos de siempre fuera de nuestro Gobierno, alejarlos del dinero público que les causa tanta tentación. Un verdadero cambio es desarrollar la economía y generar bienestar. #josécarlossánchez #Gobernador",
        "ad_creative_link_caption": "NO_TEXT",
        "ad_creative_link_description": "NO_TEXT",
        "ad_creative_link_title": "Acabemos con los corruptos!",
        "ad_delivery_start_time": "2021-02-28",
        "ad_delivery_stop_time": "NO_TEXT",
        "funding_entity": "Alejandro Gustavo Coca Cespedes",
        "impressions_min": 25000,
        "impressions_max": 29999,
        "currency": "BOB",
        "ad_url": "https://www.facebook.com/ads/library/?id=897368761011431",
        "social_media_facebook": "facebook",
        "social_media_instagram": "instagram",
        "search_terms": "José Carlos Sánchez Verazaín",
        "ad_creation_time": "2021-02-27",
        "potential_reach_min": 1000001
    },
    {
        "ad_id": "4037882719611124",
        "page_id": "101029268584854",
        "page_name": "José Carlos Sánchez Verazaín",
        "ad_snapshot_url": "https://www.facebook.com/ads/archive/render_ad/?id=4037882719611124&access_token=EA......Ts7nH10Jf",
        "ad_creative_body": "Candidato Humberto Sánchez hemos venido a buscarle a su casa de campaña para dejarle una invitación a debatir de cara a la Ciudadanía que Yo represento, ya tiene la invitación y espero su respuesta, si de verdad tiene una buena propuesta no se esconda. #CochabambaDespierta",
        "ad_creative_link_caption": "NO_TEXT",
        "ad_creative_link_description": "NO_TEXT",
        "ad_creative_link_title": "NO_TEXT",
        "ad_delivery_start_time": "2021-02-27",
        "ad_delivery_stop_time": "NO_TEXT",
        "funding_entity": "Alejandro Gustavo Coca Cespedes",
        "impressions_min": 60000,
        "impressions_max": 69999,
        "spend_min": 100,
        "spend_max": 100,
        "currency": "BOB",
        "ad_url": "https://www.facebook.com/ads/library/?id=4037882719611124",
        "social_media_facebook": "facebook",
        "social_media_instagram": "instagram",
        "search_terms": "José Carlos Sánchez Verazaín",
        "ad_creation_time": "2021-02-27",
        "potential_reach_max": 1000000,
        "potential_reach_min": 500001
    },
    {
        "ad_id": "107367994642274",
        "page_id": "101029268584854",
        "page_name": "José Carlos Sánchez Verazaín",
        "ad_snapshot_url": "https://www.facebook.com/ads/archive/render_ad/?id=107367994642274&access_token=.....2tDrPTs7nH10Jf",
        "ad_creative_body": "DENUNCIO A UNITEL POR DISCRIMINADORES... los medios de comunicación 🗞deben ser parte de la democracia, la voz de la ciudadanía, denuncia de los abusos. UNITEL se ha convertido en un instrumento del oficialismo 📢, traicionan y engañan a su propio pueblo❌.\n\n #josécarlossánchez \n#UnitelDiscriminador\n #SembrandoTiemposMejores #CochabambaDespierta",
        "ad_creative_link_caption": "NO_TEXT",
        "ad_creative_link_description": "NO_TEXT",
        "ad_creative_link_title": "Unitel contra la democracia",
        "ad_delivery_start_time": "2021-02-26",
        "ad_delivery_stop_time": "NO_TEXT",
        "funding_entity": "Alejandro Gustavo Coca Cespedes",
        "impressions_min": 50000,
        "impressions_max": 59999,
        "currency": "BOB",
        "ad_url": "https://www.facebook.com/ads/library/?id=107367994642274",
        "social_media_facebook": "facebook",
        "social_media_instagram": "instagram",
        "search_terms": "José Carlos Sánchez Verazaín",
        "ad_creation_time": "2021-02-26",
        "potential_reach_min": 1000001
    },
```


### List all facebook ads.

It uses pagination for show the acutal facebook ads.

**URL**

`/api/facebook/ads/list/by_job_id/?page_location=1&page_size=10&job_id=1`

**TYPE**

`GET`

**URL PARAMS**

* `page_location: int`
* `page_size: int`
* `job_id: int`


**RESPONSE**

A list of ads filtered by job_id

```json
[
    {
        "ad_id": "897368761011431",
        "page_id": "101029268584854",
        "page_name": "José Carlos Sánchez Verazaín",
        "ad_snapshot_url": "https://www.facebook.com/ads/archive/render_ad/?id=897368761011431&access_token=EA......0Jf",
        "ad_creative_body": "DESPIERTA COCHABAMBA! Es momento de echar a los corruptos de siempre fuera de nuestro Gobierno, alejarlos del dinero público que les causa tanta tentación. Un verdadero cambio es desarrollar la economía y generar bienestar. #josécarlossánchez #Gobernador",
        "ad_creative_link_caption": "NO_TEXT",
        "ad_creative_link_description": "NO_TEXT",
        "ad_creative_link_title": "Acabemos con los corruptos!",
        "ad_delivery_start_time": "2021-02-28",
        "ad_delivery_stop_time": "NO_TEXT",
        "funding_entity": "Alejandro Gustavo Coca Cespedes",
        "impressions_min": 25000,
        "impressions_max": 29999,
        "currency": "BOB",
        "ad_url": "https://www.facebook.com/ads/library/?id=897368761011431",
        "social_media_facebook": "facebook",
        "social_media_instagram": "instagram",
        "search_terms": "José Carlos Sánchez Verazaín",
        "ad_creation_time": "2021-02-27",
        "potential_reach_min": 1000001
    },

    {
        "ad_id": "107367994642274",
        "page_id": "101029268584854",
        "page_name": "José Carlos Sánchez Verazaín",
        "ad_snapshot_url": "https://www.facebook.com/ads/archive/render_ad/?id=107367994642274&access_token=.....2tDrPTs7nH10Jf",
        "ad_creative_body": "DENUNCIO A UNITEL POR DISCRIMINADORES... los medios de comunicación 🗞deben ser parte de la democracia, la voz de la ciudadanía, denuncia de los abusos. UNITEL se ha convertido en un instrumento del oficialismo 📢, traicionan y engañan a su propio pueblo❌.\n\n #josécarlossánchez \n#UnitelDiscriminador\n #SembrandoTiemposMejores #CochabambaDespierta",
        "ad_creative_link_caption": "NO_TEXT",
        "ad_creative_link_description": "NO_TEXT",
        "ad_creative_link_title": "Unitel contra la democracia",
        "ad_delivery_start_time": "2021-02-26",
        "ad_delivery_stop_time": "NO_TEXT",
        "funding_entity": "Alejandro Gustavo Coca Cespedes",
        "impressions_min": 50000,
        "impressions_max": 59999,
        "currency": "BOB",
        "ad_url": "https://www.facebook.com/ads/library/?id=107367994642274",
        "social_media_facebook": "facebook",
        "social_media_instagram": "instagram",
        "search_terms": "José Carlos Sánchez Verazaín",
        "ad_creation_time": "2021-02-26",
        "potential_reach_min": 1000001
    },
```


## Indicators

Once you have Jobs and Facebook ads, you can start using the `Indicators` from the Extraction, Transformation and Cleaning services.


**URL**

`/api/facebook/ads/etl/ind_a_b/:job_id`

**TYPE**

`GET`

**URL PARAMS**

* `job_id: int`

**example**

Example with job_id = 1

`/api/facebook/ads/etl/ind_a_b/1`


**RESPONSE**

Returns the fields:
* `job_id: int`,
* `ind_one_a: list({timestamp: string, count_punt: string }) `
* `ind_one_b: list({timestamp: string, count_punt: string }) `.


```json
{
    "job_id": 1,
    "ind_one_a": [
        {
            "timestamp": "2021-01-03",
            "count_punt": "1"
        },
        {
            "timestamp": "2021-01-04",
            "count_punt": "0"
        },
    ],
    "ind_one_b": [
        {
            "timestamp": "2021-01-03",
            "count_accum": "1"
        },
        {
            "timestamp": "2021-01-04",
            "count_accum": "1"
        },
        {
            "timestamp": "2021-01-05",
            "count_accum": "3"
        },
        {
            "timestamp": "2021-01-06",
            "count_accum": "3"
        },
        {
            "timestamp": "2021-01-07",
            "count_accum": "3"
        },

    ]
}
```

