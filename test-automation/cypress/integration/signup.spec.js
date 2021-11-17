/// <reference types="Cypress" />

describe('Sign up page', () => {
    const fName = 'test';
    const lName = 'tester';
    const email = 'ttester@gmail.com';
    const pwd = '1234';
    const wrongPwd = '123';
    const bio = 'I am a tester';

    beforeEach(() => {
        cy.visit('signup');

        cy.get('[data-cy="first-name-input"]').as('fNameInput');
        cy.get('[data-cy="last-name-input"]').as('lNameInput');
        cy.get('[data-cy="email-input"]').as('emailInput');
        cy.get('[data-cy="username-input"]').as('usernameInput');
        cy.get('[data-cy="pwd-input"]').as('pwdInput');
        cy.get('[data-cy="confirm-pwd-input"]').as('confirmPwdInput');
        cy.get('[data-cy="bio-input"]').as('bioInput');
        cy.get('[data-cy="back-btn"]').as('backBtn');
        cy.get('[data-cy="signup-btn"]').as('signupBtn');
    });

    it('signs up with valid info and no profile photo', () =>{
        const username = 'validTester';

        //enter valid input
        cy.get('@fNameInput').type(fName);
        cy.get('@lNameInput').type(lName);
        cy.get('@emailInput').type(email);
        cy.get('@usernameInput').type(username);
        cy.get('@pwdInput').type(pwd);
        cy.get('@confirmPwdInput').type(pwd);
        cy.get('@bioInput').type(bio);

        //click signup and check redirected to login
        cy.get('@signupBtn').click();
        cy.url().should('include', '/login')

        //enter valid login
        cy.get('[data-cy="login-username-input"]').type(username);
        cy.get('[data-cy="login-pwd-input"]').type(pwd);

        //click login and check redirected to home
        cy.get('[data-cy="login-btn"]').click();
        cy.url().should('include', '/home')
    });

    it('does not sign up with a duplicate username', () => {
        const username = 'dupTester';

        //signup with username
        cy.get('@fNameInput').type(fName);
        cy.get('@lNameInput').type(lName);
        cy.get('@emailInput').type(email);
        cy.get('@usernameInput').type(username);
        cy.get('@pwdInput').type(pwd);
        cy.get('@confirmPwdInput').type(pwd);
        cy.get('@bioInput').type(bio);
        cy.get('@signupBtn').click();

        //go back to sign up page
        cy.visit('signup');

        //try signing up with same username
        cy.get('@fNameInput').type(fName);
        cy.get('@lNameInput').type(lName);
        cy.get('@emailInput').type(email);
        cy.get('@usernameInput').type(username);
        cy.get('@pwdInput').type(wrongPwd);
        cy.get('@confirmPwdInput').type(wrongPwd);
        cy.get('@bioInput').type(bio);

        //click signup and check alert
        cy.get('@signupBtn').click();
        cy.on('window:alert',(alertText)=>{
            expect(alertText).to.match(/(That username already exists, please try again.|Incorrect username or password.)/);
        });

        //go back to log in page
        cy.visit('login');

        //enter invalid login
        cy.get('[data-cy="login-username-input"]').type(username);
        cy.get('[data-cy="login-pwd-input"]').type(wrongPwd);

        //click login and check not redirected
        cy.get('[data-cy="login-btn"]').click();
        cy.url().should('include', '/login')
    });

    it('does not sign up with mismatched passwords', () => {
        const username = 'mismatchedTester';

        //signup with mismatched passwords
        cy.get('@fNameInput').type(fName);
        cy.get('@lNameInput').type(lName);
        cy.get('@emailInput').type(email);
        cy.get('@usernameInput').type(username);
        cy.get('@pwdInput').type(pwd);
        cy.get('@confirmPwdInput').type(wrongPwd);
        cy.get('@bioInput').type(bio);

        //click signup and check alert
        cy.get('@signupBtn').click();
        cy.on('window:alert',(alertText)=>{
            expect(alertText).to.match(/(Passwords don't match.|Incorrect username or password.)/);
        });

        //go back to log in page
        cy.visit('login');

        //enter invalid login
        cy.get('[data-cy="login-username-input"]').type(username);
        cy.get('[data-cy="login-pwd-input"]').type(pwd);

        //click login and check not redirected
        cy.get('[data-cy="login-btn"]').click();
        cy.url().should('include', '/login')
    });

    it('does not sign up with empty required fields', () => {
        const username = 'emptyTester';

        //signup with just username and password
        cy.get('@usernameInput').type(username);
        cy.get('@pwdInput').type(pwd);
        cy.get('@confirmPwdInput').type(pwd);
        
        //click signup and check alert
        cy.get('@signupBtn').click();
        cy.on('window:alert',(alertText)=>{
            expect(alertText).to.match(/(Please fill out all fields.|Incorrect username or password.)/);
        });

        //go back to log in page
        cy.visit('login');

        //enter invalid login
        cy.get('[data-cy="login-username-input"]').type(username);
        cy.get('[data-cy="login-pwd-input"]').type(pwd);

        //click login and check not redirected
        cy.get('[data-cy="login-btn"]').click();
        cy.url().should('include', '/login')
    });

    it('navigates back to the login page when back is clicked', () => {
        //click back and check redirected
        cy.get('@backBtn').click();
        cy.url().should('include', '/login')
    });

    it('clears its fields when navigating away and back', () => {
        const username = 'clearTester';

        //enter valid input
        cy.get('@fNameInput').type(fName);
        cy.get('@lNameInput').type(lName);
        cy.get('@emailInput').type(email);
        cy.get('@usernameInput').type(username);
        cy.get('@pwdInput').type(pwd);
        cy.get('@confirmPwdInput').type(pwd);
        cy.get('@bioInput').type(bio);

        //click back and signup
        cy.get('@backBtn').click();
        cy.get('[data-cy="login-signup-btn"]').click();

        //check that fields are empty
        cy.get('@fNameInput').invoke('text').should('equal','');
        cy.get('@lNameInput').invoke('text').should('equal','');
        cy.get('@emailInput').invoke('text').should('equal','');
        cy.get('@usernameInput').invoke('text').should('equal','');
        cy.get('@pwdInput').invoke('text').should('equal','');
        cy.get('@confirmPwdInput').invoke('text').should('equal','');
        cy.get('@bioInput').invoke('text').should('equal','');
    });
})