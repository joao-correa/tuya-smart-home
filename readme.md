# Smart Home

Small implementation to control Tuya devices using LAN (local area network) only
*OBS: this implementation is specific for led spots and led panels on version 3.3

I did this small project to be able to easily control my led lights using Siri shortcuts

# Usage

Build your scene file (see scenes.example.json) and devices.json (see devices.example.json) and fill with you current device information (deviceId, deviceKey, deviceIp)
Then you just need keep api running and call it:

```bash
go run .

curl -X POST "http://localhost:3010/api/apply-scene?scene=turn-on"
```
