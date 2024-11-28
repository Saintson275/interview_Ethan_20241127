# React Redux Toolkit Blog App

This is a [front-end assessment at Risidio](https://github.com/Risidio/frontend_api_assement/tree/main). The assessment rquires to Fetch a list of posts from the given API endpoint and display them in a table
or a list and provide a form that allows adding a new post to the list. Each post item should display this information:
- userId
- title
- completed

However, the data fetched from the provided API endpoint([posts](https://jsonplaceholder.typicode.com/posts)) does not have a "completed" field, while another API endpoint([todos](https://jsonplaceholder.typicode.com/todos)) on JSONPlaceholder does. Since this assessment is intended to build a posts web application, I believe this should be an error in the instructions.

This project is a simple blog application built using React, Redux Toolkit, and React Router. It allows users to view a list of posts, add new posts, and navigate between different sections of the application.

![Project Logo](public/demo.png)

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Folder Structure](#folder-structure)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Approach Taken](#approach-taken)
- [Results](#results)

## Installation

To run this project locally, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Saintson275/Simple_Post_Web.git
   cd Simple_Post_Web
   
2. **Install dependencies:**:

   ```bash
   npm install
   
3. **Start the development server::**:

   ```bash
   npm run dev
   
## Usage
- Home Page: Navigate to the home page (/) to view a list of posts.
- Create Post: Navigate to /create-post and fill out the form to add a new post with a title and content.

## Folder Structure
The project structure is organized as follows:
       .


    ├── public/                 # Public assets and index.html
    ├── src/                    # Source code
    │   ├── components/         # React components
    │   ├── services/           # API services and Redux setup
    │   ├── ui/                 # Reusable UI components (e.g., buttons, cards)
    │   ├── App.tsx             # Main application component with routing
    │   └── index.tsx           # Entry point, renders the App component
    ├── package.json            # Dependencies and scripts
    └── README.md               # Project README file

## Features
- Redux Toolkit: State management with Redux Toolkit, including creating slices and defining API queries.
- React Router: Navigation and routing between different components (HomePage, AddPostForm).
- API Integration: Integration with a mock API (jsonplaceholder.typicode.com) to fetch and add posts.
- Responsive Design: Uses Tailwind CSS for responsive and mobile-first design principles.
- Error Handling: Basic error handling for API requests and form submissions.
- Dialogs and Modals: Utilizes Shadcn Alert Dialog for displaying modal dialogs for success/error messages.

## Technologies Used
- [React](https://reactjs.org)
- [Redux Toolkit](https://redux-toolkit.js.org)
- [React Router](https://reactrouter.com)
- [Axios](https://axios-http.com)
- [Tailwind CSS](https://tailwindcss.com)
- [Shadcn UI](https://ui.shadcn.com)
- [jsonplaceholder.typicode.com](https://jsonplaceholder.typicode.com) (Mock API for testing)

## Approach Taken
1. Project Setup and Initialization:

   - Clone Repository: Started by cloning the provided repository (https://github.com/Risidio/frontend_api_assement.git).
   - Initial Setup: Configured the project with TypeScript and installed necessary dependencies like Axios and Redux Toolkit.

2. API Integration:
   - Axios Configuration: Utilized Axios as the HTTP client for making requests to the public API (https://jsonplaceholder.typicode.com).
   - Redux Toolkit: Implemented Redux Toolkit for managing application state, especially using the createApi function from @reduxjs/toolkit/query to handle API          endpoints (getPosts for fetching posts, addPost for adding a new post).
  
3. Components and UI Development:
   - PostList Component: Created a component (PostList.tsx) to fetch and display a list of posts. Used Redux Toolkit's useGetPostsQuery hook to fetch posts and          displayed them in a responsive grid layout using Tailwind CSS.
   - AddPostForm Component: Implemented a form (AddPost.tsx) to add new posts. Used Redux Toolkit's useAddPostMutation hook to handle form submission, including         form state management with React's useState hooks and displaying success/error messages using an AlertDialog component.

4. Routing and Navigation:
   - React Router: Configured routes in App.tsx using react-router-dom to navigate between the home page (/) displaying the list of posts and the create post page       (/create-post) showing the form for adding new posts.

5. Error Handling and UX Improvements:
   - Error Handling: Implemented basic error handling for API requests and form submissions, displaying appropriate error messages to the user.
   - UX/UI Design: Designed a clean and responsive UI using Tailwind CSS utility classes for styling components such as cards, buttons, forms, and dialogs.

This approach focused on leveraging Redux Toolkit for efficient state management, Axios for API communication, and React for building reusable components with TypeScript for type safety and improved code maintainability. Adjustments and enhancements can be made based on specific project requirements and real-world scenarios.

## Results
1. The web application has successfully fetched a list of posts from the given API endpoint and displayed them in a list of cards.
2. Since we cannot actually change the data through the provided API, we can only verify whether a post has been successfully added by observing the status code and response data in the console. As shown in the following：
![Project Logo](public/successful.png)
![Project Logo](public/failed.png)