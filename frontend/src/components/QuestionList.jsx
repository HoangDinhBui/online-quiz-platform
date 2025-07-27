import React, { useState, useEffect } from "react";
import axios from "axios";

const QuestionList = ({ classId, onSubmit }) => {
  const [questions, setQuestions] = useState([]);
  const [answers, setAnswers] = useState({});

  useEffect(() => {
    if (classId) {
      axios;
      axios
        .get(`/api/questions?class_id=${classId}`)
        .then((response) => setQuestions(response.data))
        .catch((error) => console.error(error));
    }
  }, [classId]);

  const handleAnswerChange = (questionId, answer) => {
    setAnswers({ ...answers, [questionId]: answer });
  };

  const handleSubmit = () => {
    onSubmit(answers);
  };

  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Câu hỏi</h2>

      {!Array.isArray(questions) || questions.length === 0 ? (
        <p className="text-gray-500">Không có câu hỏi nào.</p>
      ) : (
        questions.map((q) => (
          <div key={q.question_id} className="mb-4">
            <p>{q.content}</p>
            {q.options.map((opt) => (
              <label key={opt} className="block">
                <input
                  type="radio"
                  name={q.question_id}
                  value={opt}
                  onChange={() => handleAnswerChange(q.question_id, opt)}
                />
                {opt}
              </label>
            ))}
          </div>
        ))
      )}

      <button
        onClick={handleSubmit}
        className="bg-blue-500 text-white p-2 rounded"
      >
        Nộp bài
      </button>
    </div>
  );
};

export default QuestionList;
