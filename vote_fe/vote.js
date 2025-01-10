import {decodeBase64, ethers} from "ethers";


// 输入私钥
const privateKey = "0x8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61"; // 示例私钥

// 获取钱包实例
async function getWallet(privateKey) {
    // 使用私钥创建钱包对象
    const wallet = new ethers.Wallet(privateKey);
    return wallet;
}

// 调用后端接口获取交易数据
async function fetchTransactionData() {
    const response = await fetch("http://localhost:8080/vote?index=0", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    });

    const data = await response.json();
    console.log("get data from backend", data)
    return data;
}

// 使用私钥签名交易
async function signTransaction(transactionData, wallet) {
    const data = decodeBase64(transactionData.data)
    console.log("data is ", data)
    // 创建交易对象
    const tx = {
        to: transactionData.to,
        data: data,
        nonce: transactionData.nonce,
        gasLimit: transactionData.gasLimit,
        gasPrice: transactionData.gas
    };

    // 签名交易
    const signedTx = await wallet.signTransaction(tx);
    console.log("signTransaction", signedTx)
    return signedTx;
}

// 发送签名交易到后端
async function sendSignedTransactionToBackend(signedTransaction) {
    const response = await fetch("http://localhost:8080/broadcast", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ "signed_transaction":signedTransaction }),
    });

    const result = await response.json();
    console.log("Transaction Hash:", result.transactionHash);
    return result;
}

// 主执行函数
async function main() {
    try {
        // 1. 使用私钥获取钱包
        const wallet = await getWallet(privateKey);

        // 2. 从后端获取交易数据
        const transactionData = await fetchTransactionData();
        console.log("Transaction Data:", transactionData);

        // 3. 使用私钥签名交易
        const signedTransaction = await signTransaction(transactionData, wallet);
        console.log("Signed Transaction:", signedTransaction);

        // 4. 将签名后的交易发送给后端
        const result = await sendSignedTransactionToBackend(signedTransaction);
        console.log("Transaction sent successfully:", result);
    } catch (error) {
        console.error("Error:", error);
    }
}

// 执行主函数
main();
