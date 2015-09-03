# wercker-step-webhook
A Wercker step that calls a webhook for other services to receive. 

Example usage:
```
      - releasequeue/webhook@0.0.5:
         url: "http://your.domain/your/path"
```

Will send a POST to the configured url containing a json object with the following contents.
```
{"CI" : "true",
 "WERCKER" : "true",
 "WERCKER_APPLICATION_ID" : "test_wercker_webhook_step",
 "WERCKER_APPLICATION_NAME" : "test_wercker_webhook_step",
 "WERCKER_APPLICATION_OWNER_NAME" : "owner_name",
 "WERCKER_APPLICATION_URL" : "https://app.wercker.com/#application/test_wercker_webhook_step",
 "WERCKER_BUILD_ID" : "c758b748-6e91-4837-b3f3-bc02b2cf58f0",
 "WERCKER_BUILD_URL" : "https://app.wercker.com/#build/c758b748-6e91-4837-b3f3-bc02b2cf58f0",
 "WERCKER_CACHE_DIR" : "/cache",
 "WERCKER_GIT_BRANCH" : "",
 "WERCKER_GIT_COMMIT" : "",
 "WERCKER_GIT_DOMAIN" : "",
 "WERCKER_GIT_OWNER" : "owner_name",
 "WERCKER_GIT_REPOSITORY" : "",
 "WERCKER_OUTPUT_DIR" : "/pipeline/output",
 "WERCKER_PIPELINE_DIR" : "/pipeline",
 "WERCKER_REPORT_ARTIFACTS_DIR" : "/report/webhook-c3315faf-ba5e-4067-aada-e9d2fd24cfc5/artifacts",
 "WERCKER_REPORT_DIR" : "/pipeline/report",
 "WERCKER_REPORT_MESSAGE_FILE" : "/report/webhook-c3315faf-ba5e-4067-aada-e9d2fd24cfc5/message.txt",
 "WERCKER_REPORT_NUMBERS_FILE" : "/report/webhook-c3315faf-ba5e-4067-aada-e9d2fd24cfc5/numbers.ini",
 "WERCKER_ROOT" : "/pipeline/source",
 "WERCKER_SOURCE_DIR" : "/pipeline/source",
 "WERCKER_STEP_ID" : "webhook-c3315faf-ba5e-4067-aada-e9d2fd24cfc5",
 "WERCKER_STEP_NAME" : "webhook",
 "WERCKER_STEP_OWNER" : "releasequeue",
 "WERCKER_STEP_ROOT" : "/pipeline/webhook-c3315faf-ba5e-4067-aada-e9d2fd24cfc5",
 "WERCKER_WEBHOOK_URL" : "http://your_url"
}
```
Can be used in both "build" and "deploy" pipelines and will contain the appropriate parameters.

Specifying https servers based on self signed certificates, will fail the step.

Pull requests are welcome.
