
const BACKEND_URL = "http://localhost:8080";

document.getElementById("connectWallet").addEventListener("click", async () => {
    try {
        if (typeof window.ethereum === "undefined") {
            alert("MetaMask is not installed. Please install MetaMask and try again.");
            return;
        }

        // Request account access
        const accounts = await window.ethereum.request({ method: "eth_requestAccounts" });
        const account = accounts[0];
        document.getElementById("output").textContent = `Connected account: ${account}`;
    } catch (error) {
        console.error("Error connecting to MetaMask:", error);
        document.getElementById("output").textContent = "Failed to connect to MetaMask.";
    }
});

function decodeBase64ToHex(base64String) {
    const binaryString = atob(base64String); // Base64 解码
    const hexString = Array.from(binaryString)
        .map((char) => char.charCodeAt(0).toString(16).padStart(2, '0'))
        .join('');
    return '0x' + hexString;
}

document.getElementById("vote").addEventListener("click", async () => {
    try {
        const accounts = await window.ethereum.request({ method: "eth_requestAccounts" });
        const publicKey = accounts[0];
        console.log(accounts, publicKey)
        // Fetch transaction data from backend
        const response = await fetch(`${BACKEND_URL}/vote?index=0&publickey=${publicKey}`);
        if (!response.ok) {
            throw new Error("Failed to fetch transaction data from backend.");
        }
        const transactionData = await response.json();
        const decodedData = decodeBase64ToHex(transactionData.data);
        console.log("Decoded Data:", decodedData);
        console.log(response)
        // Sign transaction with MetaMask
        const tx = {
            from: publicKey,
            to: transactionData.to,
            data: decodedData,
            // gas: `0x${parseInt(transactionData.gas).toString(16)}`,
            // gasLimit: `0x${parseInt('30000000').toString(16)}`,
            // nonce: `0x${parseInt(transactionData.nonce).toString(16)}`,
        };
        console.log(tx)

        const signedTransaction = await window.ethereum.request({
            method: "eth_sendTransaction",
            params: [tx],
        });

        document.getElementById("output").textContent = `Transaction sent: ${signedTransaction}`;
        console.log("Transaction sent successfully:", signedTransaction);
    } catch (error) {
        console.error("Error during transaction:", error);
        document.getElementById("output").textContent = "Transaction failed. "+error.toString();
    }
});