# HG8347R Golang SDK

HG8347R is an [ONT](https://en.wikipedia.org/wiki/Network_interface_device#Optical_network_terminals) device for Beijing Unicom, This project provides a control interface for Golang.

### Usage

```go
router := hg8347r.New("http://192.168.1.1", "<username>", "<password>")
devices := router.ListDevices()
```

Full API document could be found at godoc.org later.

### Tested Device

<dl>
    <dt>Model</dt>
    <dd>HG8347R</dd>
    <dt>Description</dt>
    <dd>EchoLife HG8347R EPON Terminal (PX20+/PRODUCT ID:*/CHIP:*)</dd>
    <dt>Hardware Version</dt>
    <dd>627.A</dd>
    <dt>Sofeware Version</dt>
    <dd>V3R017C10S208</dd>
    <dt>Customization Info</dt>
    <dd>BJUNICOM</dd>
</dl>
