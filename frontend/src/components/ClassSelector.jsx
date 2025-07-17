import React, { useState, useEffect } from "react";
import axios from "axios";

const ClassSelector = ({ onSelectClass }) => {
  const [classes, setClasses] = useState([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/classes")
      .then((response) => setClasses(response.data))
      .catch((error) => console.error(error));
  }, []);

  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Chọn lớp</h2>
      <select
        onChange={(e) => onSelectClass(e.target.value)}
        className="p-2 border rounded w-full md:w-64"
      >
        <option value="">Chọn lớp</option>
        {classes.map((cls) => (
          <option key={cls.class_id} value={cls.class_id}>
            {cls.Name}
          </option>
        ))}
      </select>
    </div>
  );
};

export default ClassSelector;
