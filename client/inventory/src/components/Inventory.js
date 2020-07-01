import React, { useEffect, useState } from "react";
import { makeStyles, useTheme } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import DeleteIcon from "@material-ui/icons/Delete";
import EditIcon from "@material-ui/icons/Edit";
import Button from "@material-ui/core/Button";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";

const axios = require("axios");

const useStyles = makeStyles((theme) => ({
  table: {
    minWidth: 650,
  },
  fab: {
    position: "absolute",
    right: theme.spacing(2),
  },
}));

const Inventory = (props) => {
  const classes = useStyles();
  const theme = useTheme();

  const [assets, setAssets] = useState([]);

  useEffect(() => {
    axios
      .get("/inventory/v1/assets", {})
      .then(function (response) {
        if (response.data) {
          setAssets(response.data);
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  });

  const handleOnEdit = (id) => {
    console.log(id);
  };

  const handleOnDelete = (id) => {
    console.log(id);
  };

  const handleOnAdd = () => {
    console.log("add");
  };

  return (
    <React.Fragment>
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell align="left">Date Added</TableCell>
              <TableCell align="left">Brand</TableCell>
              <TableCell align="left">Model</TableCell>
              <TableCell align="left">Serial Number</TableCell>
              <TableCell align="left">Status</TableCell>
              <TableCell align="left">Update</TableCell>
              <TableCell align="left">Delete</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {assets.map((row) => (
              <TableRow key={row.id}>
                <TableCell component="th" scope="row">
                  {row.id}
                </TableCell>
                <TableCell align="left">{row.created}</TableCell>
                <TableCell align="left">{row.brand}</TableCell>
                <TableCell align="left">{row.model}</TableCell>
                <TableCell align="left">{row.serial_number}</TableCell>
                <TableCell align="left">{row.status}</TableCell>
                <TableCell align="left">
                  <Button onClick={() => handleOnEdit(row.id)}>
                    <EditIcon />
                  </Button>
                </TableCell>
                <TableCell align="left">
                  <Button onClick={() => handleOnDelete(row.id)}>
                    <DeleteIcon />
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      <Fab
        color="primary"
        aria-label="add"
        className={classes.fab}
        onClick={handleOnAdd}
      >
        <AddIcon />
      </Fab>
    </React.Fragment>
  );
};

export default Inventory;
