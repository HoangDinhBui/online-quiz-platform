import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
  stages: [
    { duration: "1m", target: 100 }, // Tăng lên 100 người dùng trong 1 phút
    { duration: "2m", target: 100 }, // Giữ 100 người dùng trong 2 phút
    { duration: "1m", target: 500 }, // Tăng lên 500 người dùng
    { duration: "2m", target: 500 }, // Giữ 500 người dùng
    { duration: "1m", target: 0 }, // Giảm về 0
  ],
  thresholds: {
    http_req_duration: ["p(95)<500"], // 95% yêu cầu dưới 500ms
    http_req_failed: ["rate<0.01"], // Tỷ lệ lỗi dưới 1%
  },
};

const BASE_URL = "http://nginx";

export default function () {
  // Bước 1: Gọi API tạo user_id
  let userRes = http.get(`${BASE_URL}/api/generate-user-id`);
  check(userRes, { "Generate user_id status 200": (r) => r.status === 200 });

  if (userRes.status !== 200) {
    console.error("Không lấy được user_id. Body:", userRes.body);
    return;
  }

  let userId;
  try {
    userId = JSON.parse(userRes.body).user_id;
  } catch (e) {
    console.error("Lỗi parse user_id:", e, "Body:", userRes.body);
    return;
  }

  // Bước 2: Lấy danh sách lớp
  let classesRes = http.get(`${BASE_URL}/api/classes`);
  check(classesRes, { "Get classes status 200": (r) => r.status === 200 });

  if (classesRes.status !== 200) {
    console.error("Không lấy được danh sách lớp. Body:", classesRes.body);
    return;
  }

  let classId;
  try {
    let classList = JSON.parse(classesRes.body);
    if (classList.length === 0) {
      console.error("Không có lớp nào.");
      return;
    }
    classId = classList[0].class_id;
  } catch (e) {
    console.error("Lỗi parse class list:", e, "Body:", classesRes.body);
    return;
  }

  // Bước 3: Lấy câu hỏi theo lớp
  let questionsRes = http.get(`${BASE_URL}/api/questions?class_id=${classId}`);
  check(questionsRes, {
    "Get questions status 200": (r) => r.status === 200,
  });

  if (questionsRes.status !== 200) {
    console.error("Không lấy được câu hỏi. Body:", questionsRes.body);
    return;
  }

  let questions;
  try {
    questions = JSON.parse(questionsRes.body);
  } catch (e) {
    console.error("Lỗi parse questions:", e, "Body:", questionsRes.body);
    return;
  }

  // Bước 4: Nộp bài
  let answers = {};
  questions.forEach((q) => {
    if (q.options && q.options.length > 0) {
      answers[q.question_id] = q.options[0]; // chọn đáp án đầu tiên
    }
  });

  let payload = JSON.stringify({
    user_id: userId,
    class_id: classId,
    answers: answers,
  });

  let submitRes = http.post(`${BASE_URL}/api/submit`, payload, {
    headers: { "Content-Type": "application/json" },
  });

  check(submitRes, { "Submit answers status 200": (r) => r.status === 200 });

  if (submitRes.status !== 200) {
    console.error("Nộp bài thất bại. Body:", submitRes.body);
  }

  sleep(1); // Nghỉ 1 giây giữa mỗi user
}
