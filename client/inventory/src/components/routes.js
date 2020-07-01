import React from "react";

import { Router } from "@reach/router";
import Login from "./Login";
import Inventory from "./Inventory";

import Path from "./Path";

const Routes = () => {
  return (
    <Router>
      <Login path={Path.HOME} />
      <Inventory path={Path.Inventory} />
    </Router>
  );
};

export default Routes;
