import React, { useState, useEffect } from "react";
import axios from "axios";
import ClassSelector from "./components/ClassSelector";
import QuestionList from "./components/QuestionList";
import Result from "./components/Result";

const App = () => {
  const [classId, setClassId] = useState("");
  const [score, setScore] = useState(null);
  const [userId, setUserId] = useState("");

  useEffect(() => {
    axios
      .get("/api/generate-user-id")
      .then((response) => setUserId(response.data.user_id))
      .catch((error) => console.error(error));
  }, []);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost/ws");
    ws.onmessage = (event) => {
      console.log("WebSocket message:", event.data);
    };
    return () => ws.close();
  }, []);

  const handleSelectClass = (id) => {
    setClassId(id);
    setScore(null);
  };

  const handleSubmit = (answers) => {
    axios
      .post("/api/submit", {
        user_id: userId,
        class_id: classId,
        answers,
      })
      .then((response) => setScore(response.data.score))
      .catch((error) => console.error(error));
  };

  return (
    <div className="container mx-auto">
      <h1 className="text-3xl font-bold text-center my-4">
        Nền tảng thi trắc nghiệm
      </h1>
      <p>User ID: {userId}</p>
      <ClassSelector onSelectClass={handleSelectClass} />
      {classId && <QuestionList classId={classId} onSubmit={handleSubmit} />}
      {score !== null && <Result score={score} />}
    </div>
  );
};

export default App;
