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
  id: "frontendQuestionId" | "title" | "difficulty" | "tags" | "paidOnly";
  dataType: "string" | "boolean" | "number";
  label: string;
  minWidth?: number;
  align?: "right";
  format?: (value: number) => string;
};

const columns: readonly Column[] = [
  { id: "frontendQuestionId", dataType: "string", label: "#", minWidth: 170 },
  { id: "title", dataType: "string", label: "Name", minWidth: 170 },
  { id: "difficulty", dataType: "string", label: "Difficulty", minWidth: 170 },
  // { id: "tags", dataType: "string", label: "Tags", minWidth: 170 },
  // { id: "paidOnly", dataType: "boolean", label: "Premium", minWidth: 170 },
];

type Question = {
  frontendQuestionId: string;
  title: string;
  titleSlug: string;
  difficulty: string;
  tags: Array<string>;
  paidOnly: boolean;
  // solved: boolean;
};

type QuestionsTableProps = {
  questions: Array<Question>;
};

export function QuestionsTable({ questions }: QuestionsTableProps) {
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  const handleChangePage = (event: unknown, newPage: number): void => {
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
                    key={question.frontendQuestionId}
                  >
                    {columns.map((column) => {
                      const value = question[column.id];
                      console.log(column.id, value);
                      return (
                        <TableCell key={column.id} align={column.align}>
                          {/* {(() => {
                            switch (typeof value) {
                              case "boolean":
                                if (value) {
                                  return (
                                    <>
                                      <CheckCircleOutlineIcon color="success" />
                                      yes
                                    </>
                                  );
                                } else {
                                  return (
                                    <>
                                      <PanoramaFishEyeIcon color="disabled" />
                                      No
                                    </>
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
                          })()} */}
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
