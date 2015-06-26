# Fulcrum

Fulcrum lets you bring Pivotal Tracker story content into Slack in response to pasted links. The default unfurling just shows the content from the meta tags on the page. This exposes a few routes that match the URLs that stories are available from on Pivotal and either returns the story content JSON _OR_ it can send an incoming webhook to a Slack room.

Currently there is no project or user information in the messages due to the necessity of making multiple API requests to get the extra data.

## Routes

### GET

GET returns the story content JSON.

### POST

POST requests require a `room` parameter which tells Fulcrum to send a webhook to that channel. 
