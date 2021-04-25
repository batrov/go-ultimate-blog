var btcNetworkAddresses = [
    {id: "btc", name: "BTC", memo: "", address: "1BN7jvbUpe9vvfUU7VU7L1dMydVgYUirr1"},
    {id: "bep2", name: "BEP2", memo: "103482971", address: "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"},
    {id: "bep20-bsc", name:"BEP20 (BSC)", memo: "", address: "0x5d99372ba4418e9fe450b9fbd8e256bdf4cf5160"},
    {id: "erc20", name:"ERC20", memo: "", address: "0x5d99372ba4418e9fe450b9fbd8e256bdf4cf5160"},
    {id: "btc-segwit", name:"BTC (SegWit)", memo: "", address: "bc1q0xvtwjgz72ahmjj3jmymq57lk5m8aajkhdlcv5"},
]

$(document).ready(function() {
    $(".card-network").hide()

    for (i=0; i < btcNetworkAddresses.length; i++) {
        var btcButton = `<a class="btn btn-success btn-btc-network m-1" id="${btcNetworkAddresses[i].id}">${btcNetworkAddresses[i].name}</a>`

        $(".btn-btc-wrapper").append(btcButton)
    }

    registerActions()
})

function registerActions() {
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
            if (btcNetworkAddresses[i].id == id) {
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

    const copyToClipboard = str => {
        const el = document.createElement('textarea');
        el.value = str;
        document.body.appendChild(el);
        el.select();
        document.execCommand('copy');
        document.body.removeChild(el);

        Swal.fire({
            position: 'center',
            icon: 'success',
            title: 'Copied to clipboard',
            showConfirmButton: false,
            timer: 1000
          })
      };

    $(".copy-card-network-memo-value").click(function(){
        copyToClipboard($(".card-network-memo-value").text())
    })

    $(".copy-card-network-deposit-address-value").click(function(){
        copyToClipboard($(".card-network-deposit-address-value").text())
    })


}
