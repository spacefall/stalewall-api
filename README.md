
# stalewall
 An api that returns a random background from:
 - Bing [b] (image of the day)
 - Chromecast [c] (screensaver) 
 - Windows Spotlight [s]
 - Unsplash [u] (source.unsplash.com/random)
 - NASA APOD [n] (disabled by default, uses demo key)

## Usage
You can find the api at https://stalewall.vercel.app/api. You can find the supported queries below.  
A chromecast like demo is also available at https://stalewall.vercel.app/demo 

## Api queries
These queries are parsed by all providers except APOD, which is why it is disabled by default.
- `res`: Asks the api to return a picture to the requested resolution (crops the image if nc is not present).
Use 0x0 as res to get the full resolution (on unsplash will return a 16:9 image).  
Example: `?res=1920x1080`  
- `nc`: Asks the api to not crop the image.  
Example: `?res=1920x1080&nc`  
- `p`: List of providers to use.  
The query is a string of characters that represent the providers to use. (you can find the character corresponding to the provider in the list above)  
Example: `?p=bcs`  