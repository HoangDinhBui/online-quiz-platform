import React from "react";

const Result = ({ score }) => {
  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Kết quả</h2>
      <p>Điểm của bạn: {score}</p>
    </div>
  );
};

export default Result;
