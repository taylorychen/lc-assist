import { useEffect, useState } from "react";
import { Typography } from "@mui/material";
import { QuestionsTable } from "./questionsTable";

export default function Dashboard() {
  const [questions, setQuestions] = useState([]);
  const getQuestions = () => {
    const url = "http://localhost:8080/questions";
    fetch(url)
      .then((response) => response.json())
      .then((questions) => {
        console.log("got questions");
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
