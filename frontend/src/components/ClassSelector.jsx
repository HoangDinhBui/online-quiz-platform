import React, { useState, useEffect } from "react";
import axios from "axios";
import { FormControl, InputLabel, Select, MenuItem } from "@mui/material";

const ClassSelector = ({ onSelectClass }) => {
  const [classes, setClasses] = useState([]);
  const [selectedClass, setSelectedClass] = useState("");

  useEffect(() => {
    axios
      .get("/api/classes")
      .then((response) => setClasses(response.data))
      .catch((error) => {
        console.error("Lỗi gọi API /api/classes:", error);
        setClasses([]); // fallback để tránh lỗi .map nếu response null
      });
  }, []);

  const handleChange = (event) => {
    const value = event.target.value;
    setSelectedClass(value);
    onSelectClass?.(value);
  };

  return (
    <FormControl fullWidth>
      <InputLabel>Chọn lớp</InputLabel>
      <Select value={selectedClass} label="Chọn lớp" onChange={handleChange}>
        <MenuItem value="">Chọn lớp</MenuItem>
        {Array.isArray(classes) &&
          classes.map((cls) => (
            <MenuItem key={cls.class_id} value={cls.class_id}>
              {cls.name}
            </MenuItem>
          ))}
      </Select>
    </FormControl>
  );
};

export default ClassSelector;
