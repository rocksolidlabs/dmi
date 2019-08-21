# dmi
Go - DMI table decoder

```go
// DMI info
// BIOS[Vendor,Version,Release Date] System Information[Manufacturer,Product Name,Version,Serial Number,UUID]
// Base Board Information[Manufacturer,Product Name,Version,Serial Number,Asset Tag,Type]
// Chassis Information[Manufacturer,Type,Version,Serial Number,Asset Tag,Boot-up State,Power Supply State,Thermal State,Security Status,Number Of Power Cords]
dmi.Init()
dmiBios := dmi.GetBIOSInformation()
dmiSystem := dmi.GetSystemInformation()
dmiBaseboard := dmi.GetBaseboardInformation()
dmiChassis := dmi.GetChassisInformation()
```