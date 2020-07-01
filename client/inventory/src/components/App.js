import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Path from "./Path";
import Login from "./Login";
import Inventory from "./Inventory";

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path={Path.Login} component={Login} />
        <Route exact path={Path.Inventory} component={Inventory} />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
