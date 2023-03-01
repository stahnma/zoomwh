# What is this?

This is a small program designed to run on a public network that can receive webhooks from zoom for participants joining/leaving a meeting. It then translates those webhooks into webhook calls to slack to notify a specific slack channel about a participant joining or leaving. This was designed to work with long-standing/running meetings. 

## Assumptions


1. You have a long-standing zoom meeting that has a constant ID. I use a paid account, I'm not sure if you have to have one or not to do this.
1. You have access/permissions to create a zoom application.
1. You have access/permissions to create a slack webhook integration.



# Running Config

you must set

`ZOOM_SECRET` which is the `secret token` provided on your zoom application. This is how zoom knows your webhook listener is the intended target.
`ZOOM_WH_SLACK_WH_URI` is the slack webhook uri. 


# Building

Checkout the source code

    go mod tidy
    go build .


# Limitations
1. This can currently be set up for one zoom meeting (because there is no logic looking what the meeting is called or anything) and one slack webhook (thus only one channel for notifications). Neither of these would be that difficult to extend or change, but I haven't had that need yet. 
1. You can't make the POST URL path anything other than "/" right now. Again, this is an easy fix, but I haven't done it yet. Ideally you could pass in a ENV var for that. 



# Setup

## Assumptions

1. You have a long-standing zoom meeting that has a constant ID. I use a paid account, I'm not sure if you have to have one or not to do this.
1. You have access/permissions to create a zoom application.
1. You have access/permissions to create a slack webhook integration.
1. You are comfortable putting a service on a public network. (I front mine with a reverse proxy).

## Create a zoom app

go to uri
create an app
fill out form
toggle Event Subscriptions
Press + Add Event Subscriptions
You'll need a public endpoint for a listener -- which is where you'll run this app. You'll see a "validate" button. Zoom send validation payload through every so often for authentication and you need to respond to them with. The app has this built in and you just need to provide your `ZOOM_SECRET` (which is `secret token` on the zoom app configuration page).  

Press the "Add Events" button.

Click on validate. 

Save.



## Create a slack webhook integration







# LICENSE
MIT