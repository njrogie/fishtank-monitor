# Fishtank Monitor
As my wife began to accumulate reptile and fishtanks (we have lots of pets!), I envisioned a project that would allow us to monitor all of the tanks over time in a central webpage. We (as of 01/2023) have 3 betta fish tanks, 1 tetra tank, 1 ecosystem tank (with shrimp and snails), and a leopard gecko that need temperature stabilization. What a better way to ensure they are in the right conditions than to pull up a webpage that shows us a summary!

I took this opportunity to start learning some cloud development. This is a pretty ambitious first web project, since it incorporates frontend, backend, and db handling, but I've tackled small aspects of each before so I figured I would give it a shot. As an professional embedded apps & systems engineer, I have always lost motivation doing web stuff because I could never find a project that was fun. But now that that barrier is out of the way, I went to work and began coding.

## Components
### 1.) HTTP API Container - Golang (/source)
This server does all of the data handling between the database, frontend server, and microcontroller. It's packaged into a container so I can easily port it to Azure/AWS/whatever. 

### 2.) API Test Simulator - Golang (/test/simulator)
This is what I used to test our api container. You can start this program with the command arg **-r** to make real-time data points (every 30 mins) rather than test data points (every 10 seconds). This is meant to run on desktop parallel to the HTTP API.

### 3.) Microcontroller code - C++ (main\_routine)
The brains accumulating all of the data for me is an ESP32-CAM microcontroller. I do plan on adding live video feed of my Leopard Gecko to the website eventually. 

## Using this project for yourself
The main purpose of making this project open-source is to show other people how it works, but if you'd like to repurpose this project for yourself here's some tips:

- The API key can be anything you want, I'd recommend you make it a little complicated. When testing, place the key in /app/key unless you alter it otherwise. 
- More to come...

