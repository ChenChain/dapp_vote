import React from "react";

const CandidateList = ({ candidates, onVote }) => {
    return (
        <ul>
            {candidates.map((candidate) => (
                <li key={candidate.id}>
                    <strong>{candidate.name}</strong> - Votes: {candidate.voteCount.toString()}
                    <button onClick={() => onVote(candidate.id)}>Vote</button>
                </li>
            ))}
        </ul>
    );
};

export default CandidateList;
