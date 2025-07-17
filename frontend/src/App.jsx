import React, { useState } from "react";
import axios from "axios";
import ClassSelector from "./components/ClassSelector";
import QuestionList from "./components/QuestionList";
import Result from "./components/Result";

const App = () => {
  const [classId, setClassId] = useState("");
  const [score, setScore] = useState(null);

  const handleSelectClass = (id) => {
    setClassId(id);
    setScore(null);
  };

  const handleSubmit = (answers) => {
    axios
      .post("http://localhost:8080/submit", {
        user_id: "user123", // Giả định user_id tạm
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
      <ClassSelector onSelectClass={handleSelectClass} />
      {classId && <QuestionList classId={classId} onSubmit={handleSubmit} />}
      {score !== null && <Result score={score} />}
    </div>
  );
};

export default App;
