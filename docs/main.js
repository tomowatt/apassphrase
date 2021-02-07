function getBackend() {
    return 'http://get.apassphrase.me/'
}

async function getPassphrase() {
    try {
        await fetch(getBackend()).then(response => response.json()).then(data => displayPassphrase(data.passphrase));
    } catch (e) {
        displayPassphrase("Could not connect to: " + getBackend());
    }
}

function displayPassphrase(data) {
    document.getElementById("passphrase").innerHTML = data;
}
