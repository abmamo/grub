import os
import slack
from pathlib import Path
from dotenv import load_dotenv
from flask import Flask, request, Response
from slackeventsapi import SlackEventAdapter
# configure env path
env_path = Path('.') / '.env'
# load env
load_dotenv(dotenv_path=env_path)
# init client
client = slack.WebClient(token=os.environ.get("SLACK_TOKEN"))
# get bot id
BOT_ID = client.api_call("auth.test")["user_id"]

# This `app` represents your existing Flask app
app = Flask(__name__)


# An example of one of your Flask app's routes
@app.route("/")
def hello():
  return "Hello there!"

@app.route("/message-count", methods=["POST"])
def message_count():
    data = request.form
    print(data)
    return Response(), 200

# Bind the Events API route to your existing Flask app by passing the server
# instance as the last param, or with `server=app`.
slack_events_adapter = SlackEventAdapter(os.environ.get("SLACK_SIGNING_SECRET"), "/slack/events", app)


# Create an event listener for "reaction_added" events and print the emoji name
@slack_events_adapter.on("reaction_added")
def reaction_added(event_data):
  emoji = event_data["event"]["reaction"]
  print(emoji)


@slack_events_adapter.on("message")
def message(payload):
    event = payload.get("event", {})
    channel_id = event.get("channel")
    user_id = event.get("user")
    text = event.get("text")
    print(BOT_ID)
    print(user_id)
    if BOT_ID != user_id:
        client.chat_postMessage(channel="#food", text="hello %s of channel %s" % (str(user_id), str(channel_id)))

# Start the server on port 3000
if __name__ == "__main__":
  app.run(port=3000)