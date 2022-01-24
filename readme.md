# Review Digest Creator

Review Digest Creator is an app in which create reports of reviews of pre configured applications. Creates digests files in `.md`
### Setup
It'll be required to have an mongoDB database installed in your machine and also create an .env file and specify the variable `DIGEST_PATH` so the digests can be created

In order to run the app, you only need to use `go run .` in the root of the project

### APIs
The app has 2 APIs. One for creating new config for digest and other for inactivate it.

The REST APIs can be seen below.

|                 |Route                            |VERB                        |
|---------------- |---------------------------------|----------------------------|
|Create App Digest|`/api/application-digests/{appId}`|            POST            |
|Delete App Digest|`/api/application-digests/{appId}`|            DELETE          |

The body of `create` api sets the configuration for the Digests.

Below you can see an example of a valid body.

```
{
	"startDigestAt" : "2022-01-22T13:40:53.165Z",
	"pageSize" : 15, 
	"sleepTime": 24 (in hours)
}
```
The only required field is `startDigestAt`. `pageSize` has 10 as default and `sleepTime` has 24.

After confiure the api, cron jobs will be schedule to generate digests in markdown files. 

The jobs can be stopped by stopping the application.

Rerunning the app will trigger reschedule all active digests will be rescheduled for the next possible avaialability according with `sleepTime` and `startDigestAt`

### Review Digests

After have an app configured, it'll schedule jobs for it every x hours. Each job configured will query in Itunes API for new reviews, updated the internal database and create the digest with only not digested reviews in which follow the sleep time stipulated.