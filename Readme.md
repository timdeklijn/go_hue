# Go Hue

## Connect to hue

* Open hue app, go to bridge and note IP address: `192.168.178.22`.
* do a `GET` on `https://192.168.178.22/api/newdeveloper`
* Do a `POST` request to `https://192.168.178.22/api` with body:

``` json
{"devicetype":"go_hue#mb tim"}
```

* Push the link button on the Hue bridge
* Send the `POST` request again and note down the username.

``` json
[
    {
        "success": {
            "username": "73rtx0fU6CMNysLyU1QkyRf7pvGcEupIL38i982-"
        }
    }
]
```

* Any other requests to the API can be made by adding the username to the endpoint:

``` 
https://192.168.178.22/api/73rtx0fU6CMNysLyU1QkyRf7pvGcEupIL38i982-/lights
```

To access information on all lights.
