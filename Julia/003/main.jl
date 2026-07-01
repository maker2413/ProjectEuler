const limit = 600_851_475_143

function solve(to::Int)
    answer = 1

    if to % 2 == 0
        answer = 2
        to ÷= 2

        while to % 2 == 0
            to ÷= 2
        end
    end

    i = 3
    while i <= isqrt(to)
        if to % i == 0
            answer = i

            while to % i == 0
                to ÷= i
            end
        end

        i += 2
    end

    if to > answer
        return to
    end

    return answer
end

println(solve(limit))
println("--------------------------------------------------")
@time solve(limit)
