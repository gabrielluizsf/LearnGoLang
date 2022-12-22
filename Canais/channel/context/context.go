package context

func MainContext(){
  contextBackground();
  withCancelContext();
  withValueContext();
  cancelContext();
  deadlineContext();
  timeoutContext();
}
