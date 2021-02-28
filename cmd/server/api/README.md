# API DOC

## Create JOB

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
    "search_terms": "", // the search term used
    "job_id":1, // ID for created this job
    "access_token": "23qf1qf12", // Used used token
    "total_token": 1212 // the total ads found (token is bad typed)
}

```



## List JOBS

This list all the cread jobs


**URL**

`/api/facebook/jobs/list/all`

**TYPE**

`GET`

**URL PARAMS**
`page_location: int`
`page_size: int`

**Example**
`/api/facebook/jobs/list/all?page_location=1&page_size=10`


**RESPONSE**

```json

{
    "search_terms": "", // the search term used
    "job_id":1, // ID for created this job
    "access_token": "23qf1qf12", // Used used token
    "total_token": 1212 // the total ads found (token is bad typed)
}

```


