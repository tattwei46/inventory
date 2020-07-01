import React, { useState } from "react";
import TextField from "@material-ui/core/TextField";
import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import { makeStyles } from "@material-ui/core/styles";
import Path from "./Path";

const axios = require("axios");

const useStyles = makeStyles((theme) => ({
  button: {
    borderRadius: "50px",
    marginTop: "10px",
    height: "45px",
  },
}));

const Login = (props) => {
  const classes = useStyles();

  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");

  const handleOnLogin = () => {
    axios
      .post("/inventory/v1/sessions", {
        user_id: userId,
        password: password,
      })
      .then(function (response) {
        if (response.status === 200) {
          console.log("success");
          window.location = Path.Inventory;
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <Box
      display={"flex"}
      flexDirection={"column"}
      justifyContent={"center"}
      alignItems={"center"}
      flex={1}
      height={"98vh"}
    >
      <Box
        display={"flex"}
        flex={1}
        flexDirection={"column"}
        justifyContent={"center"}
      >
        <form noValidate autoComplete="off">
          <div>
            <TextField
              id="userid"
              variant="outlined"
              margin="normal"
              required
              id="outlined-required"
              label="Required"
              value={userId}
              onChange={(e) => setUserId(e.target.value)}
            />
            <br />
            <TextField
              id="password"
              variant="outlined"
              margin="normal"
              required
              label="Password"
              type="password"
              autoComplete="current-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <br />
            <Button
              className={classes.button}
              color="primary"
              fullWidth
              variant="contained"
              onClick={handleOnLogin}
            >
              Login
            </Button>
          </div>
        </form>
      </Box>
    </Box>
  );
};

export default Login;
