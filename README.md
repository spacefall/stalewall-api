# gowall
 An api that returns a random background from:
 - Bing's image of the day
 - Chromecast screensaver
 - Windows Spotlight
 
 ## Api queries
 #### Chromecast
 - `cParam`: Uses specified Google Photos parameters instead of `w0-h0`
 #### Bing
 - `bMkt`: Uses specified market instead of a random one
 - `bRes`: Uses specified photo resolution instead of using `UHD`
 - `bQlt`: Uses specified photo quality instead of `100`
 - `bH`: Resizes picture to specified height
 - `bW`: Resizes picture to specified width
 #### Spotlight
 - `sLocale`: Uses specified locale instead of a random one
 - `sPortrait`: If present, pictures are going to be in portrait instead of landscape