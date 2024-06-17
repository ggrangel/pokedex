<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">POKEDEX</h1>
</p>
<p align="center">
    <em>Unleash the Pokémon World!</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/ggrangel/pokedex?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/ggrangel/pokedex?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/ggrangel/pokedex?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/ggrangel/pokedex?style=default&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Tests](#-tests)
</details>
<hr>

##  Overview

The Pokedex project is a command-line application designed to help users explore and manage Pokémon data efficiently. It facilitates functionalities such as exploring Pokémon in specific locations, inspecting caught Pokémon for detailed information, and maintaining a comprehensive list of Pokémon in the Pokedex. By integrating with the PokeAPI, the project ensures seamless data retrieval while offering a user-friendly interface for interacting with captured Pokémon. Key features include exploring, catching, inspecting, and navigating through location areas, all aimed at enhancing the user's experience within the Pokémon ecosystem.

---

##  Features

|    |   Feature         | Description |
|----|-------------------|---------------------------------------------------------------|
| ⚙️  | **Architecture**  | The project utilizes a modular architecture with components such as exploring Pokémon in specific areas, inspecting caught Pokémon, displaying help menu, and managing a Pokedex. Key components include a REPL interface, API client, and caching module. |
| 🔩 | **Code Quality**  | The codebase maintains good quality and style with clean and well-structured code. It includes error handling for user input, API calls, and cache management. The use of interfaces and separation of concerns enhances maintainability and readability. |
| 🔌 | **Integrations**  | Key integrations include pokeapi for fetching Pokémon data, external packages for API interaction, and caching mechanisms for optimizing performance. The project seamlessly integrates external dependencies to enhance functionality and improve user experience. |
| 🧩 | **Modularity**    | The codebase exhibits high modularity and reusability with distinct modules for different functionalities like caching, API requests, and CLI commands. Modular design enables easy extension, maintenance, and customization of features without affecting other components. |
| 🧪 | **Testing**       | Testing is performed using the standard Go testing framework. Tests cover functionalities like user input processing, cache operations, Pokemon data retrieval, and command execution. The test coverage ensures code reliability and robustness. |
| ⚡️  | **Performance**   | The project demonstrates efficient performance with optimized resource usage, caching mechanisms, and API interactions. The use of caching for frequently accessed data and efficient HTTP requests enhances system responsiveness and user experience. |
| 🛡️ | **Security**      | Security measures include error handling for HTTP requests, input validation to prevent vulnerabilities, and cache management to improve data protection. Access control mechanisms ensure that only authorized users can interact with the system. |

---

##  Getting Started

###  Installation

<h4>From <code>source</code></h4>

> 1. Clone the pokedex repository:
>
> ```console
> $ git clone https://github.com/ggrangel/pokedex
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd pokedex
> ```
>
> 3. Install the dependencies:
> ```console
> $ go build -o myapp
> ```

###  Usage

<h4>From <code>source</code></h4>

> Run pokedex using the command below:
> ```console
> $ ./myapp
> ```

###  Tests

> Run the test suite using the command below:
> ```console
> $ go test
> ```

---
