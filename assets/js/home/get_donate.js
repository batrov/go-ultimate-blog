var btcNetworkAddresses = [
    {name: "btc", memo: "", address: "1BN7jvbUpe9vvfUU7VU7L1dMydVgYUirr1"},
    {name: "bep2", memo: "103482971", address: "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"},
    {name: "bep20-bsc", memo: "", address: "0x5d99372ba4418e9fe450b9fbd8e256bdf4cf5160"},
    {name: "erc20", memo: "", address: "0x5d99372ba4418e9fe450b9fbd8e256bdf4cf5160"},
    {name: "btc-segwit", memo: "", address: "bc1q0xvtwjgz72ahmjj3jmymq57lk5m8aajkhdlcv5"},
]

$(document).ready(function() {
    $(".card-network").hide()
})

$(".btn-btc-network").click(function(){
    var id = this.id

    if ($(".card-network").attr("active") == id) {
        $(".card-network").attr("active", "")
        $(".card-network").hide()
        return
    }
    $(".card-network").attr("active", id)
    $(".card-network").show()

    $(".card-network-title").text(this.text)
    
    for (i=0; i < btcNetworkAddresses.length; i++) {
        if (btcNetworkAddresses[i].name == id) {
            if (btcNetworkAddresses[i].memo == "") {
                $(".card-network-memo").hide()
            } else {
                $(".card-network-memo").show()
                $(".card-network-memo-value").text(btcNetworkAddresses[i].memo)
            }

            $(".card-network-deposit-address-value").text(btcNetworkAddresses[i].address)

            return
        }
    }

})