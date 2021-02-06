function getBackend() {
    return 'https://get-passphrase.herokuapp.com/'
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
