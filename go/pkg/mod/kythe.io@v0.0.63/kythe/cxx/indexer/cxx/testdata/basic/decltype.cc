// Checks that the result of deducing `decltype` types is recorded.
//- @x defines/binding VarX
//- VarX typed IntType
int x;
//- @x ref VarX
//- @decltype ref IntType
decltype(x) y;