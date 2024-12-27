import React, { useState } from "react";
import { ethers } from "ethers";
import votingAbi from "../abi/Voting.json";

const VotingForm = ({ contractAddress, provider }) => {
    const [candidateName, setCandidateName] = useState("");

    const addCandidate = async () => {
        if (!provider) {
            alert("Please connect your wallet!");
            return;
        }

        try {
            const signer = provider.getSigner();
            const contract = new ethers.Contract(contractAddress, votingAbi, signer);
            const tx = await contract.addCandidate(candidateName);
            await tx.wait();
            alert("Candidate added!");
            setCandidateName("");
        } catch (error) {
            console.error("Error adding candidate:", error);
        }
    };

    return (
        <div>
            <input
                type="text"
                value={candidateName}
                onChange={(e) => setCandidateName(e.target.value)}
                placeholder="Candidate Name"
            />
            <button onClick={addCandidate}>Add Candidate</button>
        </div>
    );
};

export default VotingForm;
