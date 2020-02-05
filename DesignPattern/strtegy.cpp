#include <iostream>
class TaxStrategy{
public:
    virtual double calculate(const Context& context) = 0;
    virtual ~TaxStrategy();
};//父类　各国税收的父类

class CNTax: public TaxStrategy{
public:
    virtual double calculate(const Context& context){
        //......
    }
};

class USTax:public TaxStrategy{
public:
    virtual double calculate(const Context& context){
        //......
    }
};

//可以拓展其他国家的税收
class SalesOrder{
private:
    TaxStrategy *strategy;//多态
public:
    SalesOrder(StrategyFactory* strategyFactory){//工厂模式创建各国税收对象
        strategy = strategyFactory->newStratragy();
    }
    ~SalesOrder(){
        delete this->strategy;
    }

public:
    double calculateTax(){
        //...
        ConText context;

        double val = strategy->calculate(context);
        //...
        return val;
    }
};

int main()
{
    return 0;
}

