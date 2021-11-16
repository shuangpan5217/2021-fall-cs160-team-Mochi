/// <reference types="cypress" />

import { LoginPage } from "../pages/login_page_test";

const loginPage = new LoginPage();

it("mochi note login test", function () {
    loginPage.nevigate("http://localhost:3001/login");
    loginPage.enterUsername("test");
    loginPage.enterPassword("test");
    loginPage.clickLogin();
});
