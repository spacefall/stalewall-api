
# stalewall
 An api that returns a random background from:
 - Bing's image of the day
 - Chromecast screensaver
 - Windows Spotlight

## Usage
You can find the api at https://stalewall.vercel.app/api. You can find the supported queries below.  
A chromecast home-like demo is also available at https://stalewall.vercel.app/demo 

## Api queries

- `res`: Asks the api to resize the picture to the requested resolution, while keeping aspect ratio (if possible).  
Supported on: Bing and Chromecast  
Usage: `?res=1920x1080`  
- `crop`: Asks the api to crop the image to the requested resolution, requires `res`.  
Supported on: Bing and Chromecast  
Usage: `?res=1920x1080&crop`  
- `scrop`: Asks the api to crop the image to the requested resolution, using an algorithm to keep the best part of the in the frame, requires `res` and has priority over crop.  
Supported on: Bing and Chromecast  
Usage: `?res=1920x1080&scrop`  
- `mkt`: Changes the market provided to the apis.  
Supported on: Bing and Spotlight  
Usage: `?mkt=en-US`  
- `bRes`: Asks the bing api to provide an image with the specific resolution requested, this is different from using `res`, as that downscales the "raw" photo in real time. 
Supported on: Bing
Usage: `?bRes=1920x1080`  
- `bQlt`: Asks the bing api to return a photo more or less compressed.  
0 (min) = most compressed, 100 (max) = least compressed   
Supported on: Bing   
Usage: `?bQlt=100
- `portrait`: Forces the Spotlight provider to return a portrait photo   
Supported on: Spotlight  
Usage: `?mkt=en-US`  `  