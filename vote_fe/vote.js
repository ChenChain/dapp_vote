import {decodeBase64, ethers} from "ethers";


// 输入私钥
const privateKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"; // 示例私钥
const publickey = "0xcd3B766CCDd6AE721141F452C550Ca635964ce71"; // 示例私钥
// 获取钱包实例
async function getWallet(privateKey) {
    // 使用私钥创建钱包对象
    const wallet = new ethers.Wallet(privateKey);
    return wallet;
}

// 调用后端接口获取交易数据
async function fetchTransactionData() {
    const response = await fetch("http://localhost:8080/vote?index=0&publickey=" + publickey, {
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
        gasLimit: 30000000,
        gasPrice: transactionData.gas,
        chainId: 31337,
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
