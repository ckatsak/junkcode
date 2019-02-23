#include <iostream>

class Foo
{
	public:
	Foo(int a)
	{
		std::cout << "\tFoo's constructor called with (a=" << a << \
			")" << std::endl;
	}
};

class Bar : public Foo
{
	public:
	Bar(int a, int b) : Foo(a)
	{
		std::cout << "\tBar's constructor called with (a="<< a << \
			", b=" << b << ")" << std::endl;
	}
};

class Baz : public Bar
{
	public:
	Baz(int a, int b, int c) : Bar(a, b)
	{
		std::cout << "\tBaz's constructor called with (a=" << a << \
			", b=" << b << ", c=" << c << ")" << std::endl;
	}
};

int main(void)
{
	std::cout << "New Bar:" << std::endl;
	Bar bar(1, 2);

	std::cout << "New Baz:" << std::endl;
	Baz baz(3, 2, 1);

	return 0;
}
