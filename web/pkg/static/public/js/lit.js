import {
  LitElement,
  html,
} from "https://cdn.jsdelivr.net/gh/lit/dist@3/core/lit-core.min.js";

export class MyElement extends LitElement {
  static properties = {
    version: {},
  };

  constructor() {
    super();
    this.version = "STARTING";
  }

  render() {
    return html`
      <p>Welcome to the Lit tutorial!</p>
      <p>This is the ${this.version} code.</p>
    `;
  }
}
customElements.define("my-element", MyElement);
