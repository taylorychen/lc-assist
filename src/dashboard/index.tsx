import { useEffect, useState } from "react";
import { Typography } from "@mui/material";
import { QuestionsTable } from "./questionsTable";

export default function Dashboard() {
  const [questions, setQuestions] = useState([]);
  const getQuestions = () => {
    fetch(`/api/questions?amount=10`)
      .then((response) => response.json())
      .then((questions) => {
        setQuestions(questions);
      });
  };

  useEffect(getQuestions, []);

  return (
    <>
      <Typography variant="h3">Username</Typography>
      <QuestionsTable questions={questions} />
    </>
  );
}
