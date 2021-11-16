export class LoginPage {
    loginPage_username = ":nth-child(2) > .agenda";
    loginPage_password = ":nth-child(3) > .agenda";
    loginPage_loginButton = ".primary";

    nevigate(url) {
        cy.visit(url);
    }
    enterUsername(username) {
        cy.get(this.loginPage_username).type(username);
    }
    enterPassword(password) {
        cy.get(this.loginPage_password).type(password);
    }
    clickLogin() {
        cy.get(this.loginPage_loginButton).click();
    }
}
