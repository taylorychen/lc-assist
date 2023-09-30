import { useState } from "react";
import CheckCircleOutlineIcon from "@mui/icons-material/CheckCircleOutline";
import PanoramaFishEyeIcon from "@mui/icons-material/PanoramaFishEye";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TablePagination from "@mui/material/TablePagination";
import TableRow from "@mui/material/TableRow";

type Column = {
  id: "name" | "number" | "difficulty" | "type" | "solved";
  dataType: "string" | "boolean" | "number";
  label: string;
  minWidth?: number;
  align?: "right";
  format?: (value: number) => string;
};

const columns: readonly Column[] = [
  { id: "number", dataType: "string", label: "#", minWidth: 170 },
  { id: "name", dataType: "string", label: "Name", minWidth: 170 },
  { id: "difficulty", dataType: "string", label: "Difficulty", minWidth: 170 },
  { id: "type", dataType: "string", label: "Type", minWidth: 170 },
  { id: "solved", dataType: "boolean", label: "Solved", minWidth: 170 },
];

type Question = {
  number: number;
  name: string;
  difficulty: string;
  type: string;
  solved: boolean;
};

// function createData(
//   number: number,
//   name: string,
//   difficulty: string,
//   type: string,
//   solved: boolean
// ): Question {
//   return { number, name, difficulty, type, solved };
// }

// const rows = [
//   createData(1, "BST", "E", "Tree"),
//   createData(2, "Trie", "E", "Trie"),
//   createData(3, "Dict", "M", "Hashmap"),
//   createData(4, "Linked List", "M", "Two pointer"),
//   createData(5, "Enumerate", "M", "Python"),
//   createData(6, "Queue", "M", "Queue"),
//   createData(7, "Breadth-First-Search", "M", "Graph"),
//   createData(8, "Kruskal's Algorithm", "H", "Graph"),
//   createData(9, "Binary Tree", "E", "Tree"),
//   createData(10, "Prim's Algorithm", "M", "Graph"),
//   createData(11, "Dijkstra", "M", "Graph"),
//   createData(12, "Hello World", "E", "String"),
// ];

type QuestionsTableProps = {
  questions: Array<Question>;
};

export function QuestionsTable({ questions }: QuestionsTableProps) {
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  const handleChangePage = (event: unknown, newPage: number) => {
    console.log(event);
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  return (
    <Paper sx={{ width: "100%", overflow: "hidden" }}>
      <TableContainer sx={{ maxHeight: 440 }}>
        <Table stickyHeader>
          <TableHead>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth }}
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {questions
              .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
              .map((question) => {
                return (
                  <TableRow
                    hover
                    role="checkbox"
                    tabIndex={-1}
                    key={question.number}
                  >
                    {columns.map((column) => {
                      const value = question[column.id];
                      return (
                        <TableCell key={column.id} align={column.align}>
                          {(() => {
                            switch (typeof value) {
                              case "boolean":
                                if (value) {
                                  return (
                                    <CheckCircleOutlineIcon color="success" />
                                  );
                                } else {
                                  return (
                                    <PanoramaFishEyeIcon color="disabled" />
                                  );
                                }
                              case "string":
                              case "number":
                              case "bigint":
                              case "symbol":
                              case "undefined":
                              case "object":
                              case "function":
                                return null;
                            }
                          })()}
                          {column.format && typeof value === "number"
                            ? column.format(value)
                            : value}
                        </TableCell>
                      );
                    })}
                  </TableRow>
                );
              })}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 25, 100]}
        component="div"
        count={questions.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
      />
    </Paper>
  );
}
