#import <Foundation/Foundation.h>

int main (int argc, const char * argv[]) {
    NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init];
    int a,b;
    scanf("%d %d", &a, &b);
    printf("%d", a+b);
    [pool drain];
    return 0;
}