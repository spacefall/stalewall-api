
# stalewall
 An api that returns a random background from:
 - Bing's image of the day
 - Chromecast screensaver
 - Windows Spotlight
 
 ## Api queries

 #### Bing
 - `bMkt`: Uses specified market instead of a random one
 - `bRes`: Uses specified photo resolution instead of using `UHD`
 - `bQlt`: Uses specified photo quality instead of `100`

 #### Spotlight
 - `sLocale`: Uses specified locale instead of a random one
 - `sPortrait`: If present, pictures are going to be in portrait instead of landscape

#### Common
- `res`: Tries to resize picture to requested resolution, while keeping aspect ratio. It must be written like this: `width`x`height`, e.g. `1920x1080`. This query is valid only for bing and chromecast and only if the `raw` query isn't specified.
- `crop`: Crops the image from the center, requires `res`, is disabled when `raw` or `scrop` are specified. Works with chromecast and bing.
- `scrop`: Aka smart crop, crops using an algorithm that finds the interesting part of the image, requires `res`, is disabled when `raw` is specified. Works with chromecast and bing.
