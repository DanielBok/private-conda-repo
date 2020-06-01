import { push } from "connected-react-router";
import React, { Component } from "react";
import { connect } from "react-redux";
import { RouteComponentProps, withRouter } from "react-router";
import { Dispatch } from "redux";
import ErrorPage from "./ErrorPage";

type State = {
  hasError: boolean;
};

type Props = RouteComponentProps & {
  errorRoute?: string;
  goto: (path: string) => void;
  children: React.ReactNode;
};

class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      hasError: false,
    };

    props.history.listen(() => {
      if (this.state.hasError) {
        this.setState({ hasError: false });
      }
    });
  }

  static getDerivedStateFromError(error: Error) {
    return { hasError: typeof error === "object" };
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo): void {
    if (process.env.NODE_ENV === "development") {
      console.error(error);
      console.error(errorInfo);
    }

    if (typeof this.props.errorRoute === "string") {
      this.props.goto(this.props.errorRoute);
    }
  }

  render() {
    if (this.state.hasError) {
      return <ErrorPage />;
    }
    return this.props.children;
  }
}

const mapDispatchToProps = (dispatch: Dispatch) => ({
  goto: (path: string) => dispatch(push(path)),
});

export default connect(null, mapDispatchToProps)(withRouter(ErrorBoundary));
