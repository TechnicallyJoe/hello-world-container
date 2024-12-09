import azure.functions as func
import datetime
import json
import logging

app = func.FunctionApp()

@app.route(route="httptrigger", auth_level=func.AuthLevel.ANONYMOUS)
def httptrigger(req: func.HttpRequest) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')

    inputString = req.params.get('inputString')
    if not inputString:
        try:
            req_body = req.get_json()
        except ValueError:
            pass
        else:
            inputString = req_body.get('inputString')

    if inputString:
        return func.HttpResponse(f"{inputString}")
    else:
        return func.HttpResponse(
             "This HTTP triggered function executed successfully. Pass a name in the query string or in the request body for a personalized response.",
             status_code=200
        )