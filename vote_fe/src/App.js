import React, { useState, useEffect } from "react";
import { ethers } from "ethers";
import CandidateList from "./components/CandidateList";
import VotingForm from "./components/VotingForm";
import votingAbi from "./abi/Voting.json";

const CONTRACT_ADDRESS = "YOUR_CONTRACT_ADDRESS";

function App() {
  const [candidates, setCandidates] = useState([]);
  const [isConnected, setIsConnected] = useState(false);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);

  useEffect(() => {
    loadCandidates();
  }, []);

  const loadCandidates = async () => {
    try {
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const contract = new ethers.Contract(CONTRACT_ADDRESS, votingAbi, provider);
      const data = await contract.getCandidates();
      setCandidates(data.map((c, index) => ({ ...c, id: index })));
    } catch (error) {
      console.error("Error loading candidates:", error);
    }
  };

  const connectWallet = async () => {
    try {
      await window.ethereum.request({ method: "eth_requestAccounts" });
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();
      setProvider(provider);
      setSigner(signer);
      setIsConnected(true);
    } catch (error) {
      console.error("Error connecting wallet:", error);
    }
  };

  const vote = async (candidateIndex) => {
    try {
      if (!signer) {
        alert("Please connect your wallet first!");
        return;
      }
      const contract = new ethers.Contract(CONTRACT_ADDRESS, votingAbi, signer);
      const tx = await contract.vote(candidateIndex);
      await tx.wait();
      alert("Vote successful!");
      loadCandidates();
    } catch (error) {
      console.error("Error voting:", error);
    }
  };

  return (
      <div>
        <h1>Decentralized Voting DApp</h1>
        {!isConnected ? (
            <button onClick={connectWallet}>Connect Wallet</button>
        ) : (
            <div>
              <h2>Candidates</h2>
              <CandidateList candidates={candidates} onVote={vote} />
            </div>
        )}
      </div>
  );
}

export default App;
