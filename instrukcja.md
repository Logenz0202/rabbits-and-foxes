# Symulacja Lisy i Króliki (i trawa)

To klasyczne zadanie jest ostatnim, które sugeruję abyście napisali. Pracujemy już nad nim od tygodnia / dwóch, i po oddaniu Lasu, będzie to pewnie dość ciekawe do napisania. Wymagania co do tego zadania są następujące:

1. Napisz to w grafice. Możesz wykorzystać Raylib, SDL2, Fyne, Ebiten lub inne popularne rozwiązania.
2. Symulacja powinna zakładać współistnienie 3 elementów: trawy, królików i lisów.
3. Trawa może pojawić się na pustych polach i rosnąć (jej ilość zmienia się od 0, do zadanego max).
4. Króliki mogą przemieszczać się po planszy (losowo, lub mogą "widzieć" otoczenie kilku sąsiednich pól).
5. Królik je trawę (która wtedy znika lub zmniejsza się jej ilość). Najedzony królik, który spotka innego najedzonego królika, rozmnaża się, i powstaje trzeci królik (od razu doroły). Króliki "rozmnożone" przez jakiś czas się nie reprodukują.
6. Na planszy grasują lisy, które nie jedzą trawy. Jedzą króliki. Lisy mogą poruszać się losowo lub "widzieć"; lis, który zje królika jest najedzony jakiś czas i po spotkaniu drugiego najedzonego lista, mogą "zrobić" kolejnego lisa.
7. Energia zwierząt spada, jeżeli nie jedzą. Gdy osiągnie 0, następuje śmierć i zwierzę "znika".
8. Pozostałe ważne z punktu widzenia symulacji parametry oraz sposób rozwiązania, ustal samodzielnie.
9. W zadaniu należy przeprowadzić symulację w której da się obserwować liczbę lisów i królików oraz dynamikę zmian w ich populacji.
10. Program powinien działać pod Linux bez żadnych specjalnych zabiegów - zatem sprawdź np. w wirtualce, czy nie będzie jakichś problemów. Sprawdzać go będę pod Linux - zatem rozwiązania typowo 'windowe' odpadają.

W idealnym przypadku Twój program może wykorzystywać Fyne, aby pozwolić na początkowe ustawienie gęstości populacji - liczby królików, lisów, prędkości wzrostu trawy, i innych parametrów. Następnie w oknie graficznym program powinien rysować planszę na której przestawiona jest bieżąca sytuacja świata. Nie musisz rysować królików/lisów ładnie - wystarczy kropka/piksel. W dolnej części ekranu powinien znaleźć się wykres aktualizowany na bieżąco z liczbami królików, lisów i trawy, który pokaże dynamicznie zmiany w populacji.

Uproszczony mockup projektu pokazuję na rysunku - widać w nim tło z królikami i lisami. W praktyce mogą to być nawet pojedyncze piksele (dla większych rozmiarów planszy). Warto umieścić wykres i przyciski których kliknięcie uruchomi lub zatrzyma symulację. Możesz dodać "rysowanie" myszą pozycji w których będą króliki i lisy - aby jakoś je rozmieścić. Aby lepiej zrozumieć jaka jest dynamika takiego systemu oraz czym jest stała reprodukcji, obejrzyj film z Veritasium, który podawałem wcześniej.

### Do oddania:

Program w wersji źródłowej.
Program w wersji zoptymalizowanej, skompilowanej pod Linux
Dokumentacja: opracowanie na temat działania algorytmów, które wykorzystałeś, ciekawe miejsca, Twoja ocena i opinia o symulacji, co odkryłeś lub odkryłaś, itp.
