const {ethers} = require("hardhat");

const provider = new ethers.providers.JsonRpcProvider();
const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"; // 合约地址
const abi = "./artifacts/contracts/Audi.sol/Audi.json"; // 合约的 ABI，可从编译后的文件或构建文件夹中获取
const audiContract = new ethers.Contract(contractAddress, abi, provider);
audiContract.audiTotalSupply().then((supply) => console.log("总供应量：", supply));