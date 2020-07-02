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
import SearchIcon from "@material-ui/icons/Search";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import TextField from "@material-ui/core/TextField";
import DateFnsUtils from "@date-io/date-fns";
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from "@material-ui/pickers";
import Path from "./Path";
import moment from "moment";
import InputLabel from "@material-ui/core/InputLabel";
import MenuItem from "@material-ui/core/MenuItem";
import FormControl from "@material-ui/core/FormControl";
import Select from "@material-ui/core/Select";

const axios = require("axios");

const useStyles = makeStyles((theme) => ({
  root: {
    "& > *": {
      margin: theme.spacing(1),
    },
  },
  table: {
    minWidth: 650,
  },
  fab: {
    position: "absolute",
    right: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
}));

const Inventory = (props) => {
  const classes = useStyles();

  const [assets, setAssets] = useState([]);
  const [open, setOpen] = useState(false);
  const [search, setSearch] = useState(false);
  const [update, setUpdate] = useState(false);

  const [id, setId] = useState("");
  const [serialNumber, setSerialNumber] = useState("");
  const [brand, setBrand] = useState("");
  const [model, setModel] = useState("");
  const [status, setStatus] = useState("");

  const handleStatusChange = (event) => {
    setStatus(event.target.value);
  };

  const [selectedDate, setSelectedDate] = React.useState(
    new Date("2020-07-01T00:00:00")
  );

  const [selectedFrom, setSelectedFrom] = React.useState(
    new Date("2020-07-01T00:00:00")
  );

  const [selectedTo, setSelectedTo] = React.useState(
    new Date("2020-07-01T00:00:00")
  );

  useEffect(() => {
    axios
      .get("/inventory/v1/assets")
      .then(function (response) {
        if (response.data) {
          setAssets(response.data);
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  }, [open]);

  const handleOnDateChange = (date) => {
    setSelectedDate(date);
  };

  const handleOnFromChange = (date) => {
    setSelectedFrom(date);
  };

  const handleOnToChange = (date) => {
    setSelectedTo(date);
  };

  const handleOnOpen = () => {
    setOpen(true);
  };

  const handleOnSearchOpen = () => {
    setSearch(true);
  };

  const handleOnClose = () => {
    if (update) setUpdate(false);
    setOpen(false);
  };

  const handleOnSearchClose = () => {
    setSearch(false);
  };

  const handleOnSearch = () => {
    var range = {
      from: moment(selectedFrom).format("YYYY-MM-DD"),
      to: moment(selectedTo).format("YYYY-MM-DD"),
    };

    axios
      .post("/inventory/v1/assets/action/search", {
        range: range,
        model: model,
        brand: brand,
        serial_number: serialNumber,
        status: status,
      })
      .then(function (response) {
        if (response.status === 200) {
          setAssets(response.data);
        }
        if (response.status === 204) {
          setAssets([]);
        }

        setSearch(false);
      })
      .catch(function (error) {
        console.log(error);
        setSearch(false);
      });
  };

  const handleOnAdd = () => {
    if (update) {
      axios
        .put("/inventory/v1/assets/" + id, {
          created: moment(selectedDate).format("YYYY-MM-DD"),
          model: model,
          brand: brand,
          serial_number: serialNumber,
          status: status,
        })
        .then(function (response) {
          if (response.status === 200) {
            setUpdate(false);
            setOpen(false);
          }
        })
        .catch(function (error) {
          console.log(error);
          setUpdate(false);
          setOpen(false);
        });
    } else {
      axios
        .post("/inventory/v1/assets", {
          created: selectedDate,
          model: model,
          brand: brand,
          serial_number: serialNumber,
        })
        .then(function (response) {
          if (response.status === 200) {
            setOpen(false);
          }
        })
        .catch(function (error) {
          console.log(error);
          setOpen(false);
        });
    }
  };

  const handleOnDelete = (id) => {
    axios
      .delete("/inventory/v1/assets/" + id)
      .then(function (response) {
        if (response.status === 200) {
          // TODO : HOW TO HANDLE IT MORE GRACEFULLY
          window.location = Path.Inventory;
        }
      })
      .catch(function (error) {
        console.log(error);
        // TODO : HOW TO HANDLE IT MORE GRACEFULLY
        window.location = Path.Inventory;
      });
  };

  const handleOnEdit = (asset) => {
    setId(asset.id);
    setSelectedDate(asset.created);
    setBrand(asset.brand);
    setModel(asset.model);
    setStatus(asset.status);
    setSerialNumber(asset.serial_number);

    setUpdate(true);
    setOpen(true);
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
                  <Button onClick={() => handleOnEdit(row)}>
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

      <div className={classes.fab}>
        {/*Floating Action Button to Add Items*/}
        <Fab color="primary" aria-label="add" onClick={handleOnOpen}>
          <AddIcon />
        </Fab>
        {/*Floating Action Button to Search Items*/}
        <Fab color="secondary" aria-label="edit" onClick={handleOnSearchOpen}>
          <SearchIcon />
        </Fab>
      </div>

      {/*Form Dialog to Add Items*/}
      <Dialog
        open={open}
        onClose={handleOnClose}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">
          {update ? "Update" : "Add"} Items
        </DialogTitle>
        <DialogContent>
          <div>
            {update ? (
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDatePicker
                  disableToolbar
                  variant="inline"
                  format="MM/dd/yyyy"
                  margin="normal"
                  id="date-picker-inline"
                  label="Date"
                  value={selectedDate}
                  onChange={handleOnDateChange}
                  KeyboardButtonProps={{
                    "aria-label": "change date",
                  }}
                />
              </MuiPickersUtilsProvider>
            ) : null}
          </div>

          <TextField
            required
            autoFocus
            id="brand"
            label="Brand"
            type="text"
            fullWidth
            value={update ? brand : null}
            onChange={(e) => setBrand(e.target.value)}
          />
          <TextField
            required
            autoFocus
            id="model"
            label="Model"
            type="text"
            fullWidth
            value={update ? model : null}
            onChange={(e) => setModel(e.target.value)}
          />
          <TextField
            required
            autoFocus
            id="serial_number"
            label="Serial Number"
            type="text"
            fullWidth
            value={update ? serialNumber : null}
            onChange={(e) => setSerialNumber(e.target.value)}
          />
          {update ? (
            <FormControl className={classes.formControl}>
              <InputLabel id="demo-simple-select-label">Status</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={update ? status : null}
                onChange={handleStatusChange}
              >
                <MenuItem value={"Available"}>Available</MenuItem>
                <MenuItem value={"Not Available"}>Not Available</MenuItem>
              </Select>
            </FormControl>
          ) : null}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleOnClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleOnAdd} color="primary">
            {update ? `Update` : `Add`}
          </Button>
        </DialogActions>
      </Dialog>

      {/*Form Dialog to Search Items*/}
      <Dialog
        open={search}
        onClose={handleOnSearchClose}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title-search">Search Items</DialogTitle>
        <DialogContent>
          <div>
            <MuiPickersUtilsProvider utils={DateFnsUtils}>
              <KeyboardDatePicker
                disableToolbar
                variant="inline"
                format="MM/dd/yyyy"
                margin="normal"
                id="date-picker-inline"
                label="From"
                value={selectedFrom}
                onChange={handleOnFromChange}
                KeyboardButtonProps={{
                  "aria-label": "change date",
                }}
              />
            </MuiPickersUtilsProvider>

            <MuiPickersUtilsProvider utils={DateFnsUtils}>
              <KeyboardDatePicker
                disableToolbar
                variant="inline"
                format="MM/dd/yyyy"
                margin="normal"
                id="date-picker-inline"
                label="To"
                value={selectedTo}
                onChange={handleOnToChange}
                KeyboardButtonProps={{
                  "aria-label": "change date",
                }}
              />
            </MuiPickersUtilsProvider>
          </div>

          <TextField
            required
            autoFocus
            id="brand"
            label="Brand"
            type="text"
            fullWidth
            value={update ? brand : null}
            onChange={(e) => setBrand(e.target.value)}
          />
          <TextField
            required
            autoFocus
            id="model"
            label="Model"
            type="text"
            fullWidth
            value={update ? model : null}
            onChange={(e) => setModel(e.target.value)}
          />
          <TextField
            required
            autoFocus
            id="serial_number"
            label="Serial Number"
            type="text"
            fullWidth
            value={update ? serialNumber : null}
            onChange={(e) => setSerialNumber(e.target.value)}
          />

          <FormControl className={classes.formControl}>
            <InputLabel id="demo-simple-select-label">Status</InputLabel>
            <Select
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={status}
              onChange={handleStatusChange}
            >
              <MenuItem value={"Available"}>Available</MenuItem>
              <MenuItem value={"Not Available"}>Not Available</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleOnSearchClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleOnSearch} color="primary">
            Search
          </Button>
        </DialogActions>
      </Dialog>
    </React.Fragment>
  );
};

export default Inventory;
